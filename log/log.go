package log

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/isomnath/belvedere/config"

	"github.com/sirupsen/logrus"
	ddTracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const (
	postgres = "log.postgres.query"
	mongo    = "log.mongo.query"
	redis    = "log.redis.query"
)

const (
	na          = "log.na"
	httpRequest = "log.http.request"
	datastore   = "log.datastore"
)

var persistenceTypes map[string]string

type logSpan struct {
	span  tracer.Span
	isNew bool
}

type Logger struct {
	*logrus.Logger
}

var Log *Logger

type ErrorLogger struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		panic(ErrorLogger{Error: err})
	}
}

func Setup() {
	level, err := logrus.ParseLevel(config.GetAppLogLevel())
	panicIfError(err)

	persistenceTypes = map[string]string{
		postgres: "PostgreSQL",
		mongo:    "MongoDB",
		redis:    "Redis",
	}

	logrusVars := &logrus.Logger{
		Out:       os.Stderr,
		Hooks:     make(logrus.LevelHooks),
		Formatter: &logrus.JSONFormatter{},
		Level:     level,
	}
	logrusVars.AddHook(&ddTracer.DDContextLogHook{})

	Log = &Logger{logrusVars}
}

func (logger *Logger) getBaseLogEntry() *logrus.Entry {
	return logger.WithFields(
		logrus.Fields{
			"application": map[string]string{
				"name":        config.GetAppName(),
				"version":     config.GetAppVersion(),
				"environment": config.GetAppEnvironment(),
			},
		})
}

func (logger *Logger) baseLogEntry() *logrus.Entry {
	return logger.getBaseLogEntry().WithField("context", na)
}

func (logger *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.baseLogEntry().
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Errorf(format, args...)
	ls.finish()
}

func (logger *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.baseLogEntry().
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Errorf(format, args...)
	ls.finish()
}

func (logger *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.baseLogEntry().
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Infof(format, args...)
	ls.finish()
}

func (logger *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.baseLogEntry().
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Warnf(format, args...)
	ls.finish()
}

func (logger *Logger) httpRequestLogEntry(r *http.Request) *logrus.Entry {
	return logger.getBaseLogEntry().
		WithFields(logrus.Fields{
			"context": httpRequest,
			"request": map[string]interface{}{
				"path":           r.URL.Path,
				"method":         r.Method,
				"host":           r.Host,
				"remote_address": r.RemoteAddr,
			},
		})
}

func (logger *Logger) HTTPStatInfo(r *http.Request, startTime, responseTime time.Time, statusCode int) {
	latency := responseTime.Sub(startTime)
	fields := logrus.Fields{
		"context": httpRequest,
		"request": map[string]interface{}{
			"host":           r.Host,
			"method":         r.Method,
			"path":           r.URL.Path,
			"start_time":     startTime.Format(time.RFC3339),
			"remote_address": r.RemoteAddr,
		},
		"response": map[string]interface{}{
			"end_time": responseTime.Format(time.RFC3339),
			"latency":  fmt.Sprintf("%d ms", latency.Milliseconds()),
			"status":   statusCode,
		},
		"forwarded_headers": r.Header.Get("X_FORWARDED-FOR"),
	}
	ls := logger.spanFromContext(r.Context())
	logger.httpRequestLogEntry(r).
		WithFields(fields).
		WithFields(logger.getTracerSpanFields(ls.span)).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		Infof("http log")
	ls.finish()
}

func (logger *Logger) HTTPErrorf(r *http.Request, format string, args ...interface{}) {
	ls := logger.spanFromContext(r.Context())
	logger.httpRequestLogEntry(r).
		WithFields(logger.getTracerSpanFields(ls.span)).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		Errorf(format, args...)
	ls.finish()
}

func (logger *Logger) HTTPInfof(r *http.Request, format string, args ...interface{}) {
	ls := logger.spanFromContext(r.Context())
	logger.httpRequestLogEntry(r).
		WithFields(logger.getTracerSpanFields(ls.span)).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		Infof(format, args...)
	ls.finish()
}

func (logger *Logger) HTTPWarnf(r *http.Request, format string, args ...interface{}) {
	ls := logger.spanFromContext(r.Context())
	logger.httpRequestLogEntry(r).
		WithFields(logger.getTracerSpanFields(ls.span)).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		Warnf(format, args...)
	ls.finish()
}

func (logger *Logger) persistenceLogEntry(dbType string) *logrus.Entry {
	persistenceType := persistenceTypes[dbType]
	return logger.getBaseLogEntry().
		WithFields(logrus.Fields{
			"context": datastore,
			"type":    persistenceType,
		})
}

func (logger *Logger) PostgresErrorf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(postgres).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Errorf(format, args...)
	ls.finish()
}

func (logger *Logger) PostgresInfof(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(postgres).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Infof(format, args...)
	ls.finish()
}

func (logger *Logger) PostgresWarnf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(postgres).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Warnf(format, args...)
	ls.finish()
}

func (logger *Logger) MongoErrorf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(mongo).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Errorf(format, args...)
	ls.finish()
}

func (logger *Logger) MongoInfof(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(mongo).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Infof(format, args...)
	ls.finish()
}

func (logger *Logger) MongoWarnf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(mongo).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Warnf(format, args...)
	ls.finish()
}

func (logger *Logger) RedisErrorf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(redis).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Errorf(format, args...)
	ls.finish()
}

func (logger *Logger) RedisInfof(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(redis).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Infof(format, args...)
	ls.finish()
}

func (logger *Logger) RedisWarnf(ctx context.Context, format string, args ...interface{}) {
	ls := logger.spanFromContext(ctx)
	logger.persistenceLogEntry(redis).
		WithContext(ctx).
		WithFields(logger.getProcessFields(runtime.Caller(1))).
		WithFields(logger.getTracerSpanFields(ls.span)).
		Warnf(format, args...)
	ls.finish()
}

func (logger *Logger) getProcessFields(pc uintptr, file string, line int, ok bool) logrus.Fields {
	var fileName, fn string
	if !ok {
		fileName = "unknown"
		fn = "unknown"
	} else {
		fileName = file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
		fnName := runtime.FuncForPC(pc).Name()
		fn = fnName[strings.LastIndex(fnName, ".")+1:]
	}

	return logrus.Fields{
		"file":     fileName,
		"function": fn,
	}
}

func (logger *Logger) spanFromContext(ctx context.Context) logSpan {
	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		span = tracer.StartSpan(na)
		return logSpan{
			span:  span,
			isNew: true,
		}
	}
	return logSpan{
		span:  span,
		isNew: false,
	}
}

func (ls *logSpan) finish() {
	if ls.isNew {
		ls.span.Finish()
	}
}

func (logger *Logger) getTracerSpanFields(span tracer.Span) logrus.Fields {
	return logrus.Fields{
		"dd.span_id":  span.Context().SpanID(),
		"dd.trace_id": span.Context().TraceID(),
	}
}
