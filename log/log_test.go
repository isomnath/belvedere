package log

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/isomnath/belvedere/instrumentation"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"github.com/isomnath/belvedere/config"
)

type LoggerTestSuite struct {
	suite.Suite
}

func (suite *LoggerTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	instrumentation.StartDDTracer()
}

func (suite *LoggerTestSuite) TearDownSuite() {
	instrumentation.StopDDTracer()
}

func (suite *LoggerTestSuite) TestLoggerPanic() {
	os.Setenv("APP_LOG_LEVEL", "INVALID")
	config.LoadBaseConfig()
	suite.Panics(func() {
		Setup()
	})
}

func (suite *LoggerTestSuite) TestLogger() {
	Setup()
	ctx := context.Background()
	_, ctx = tracer.StartSpanFromContext(ctx, "test")

	Log.Fatalf(ctx, "test message, args1: %s", "123")
	Log.Errorf(ctx, "test message, args1: %s", "123")
	Log.Warnf(ctx, "test message, args1: %s", "123")
	Log.Infof(ctx, "test message, args1: %s", "123")

	request, _ := http.NewRequest(http.MethodGet, "/test/path/123", nil)
	startTime := time.Now()
	responseTime := startTime.Add(20 * time.Millisecond)
	Log.HTTPStatInfo(request, startTime, responseTime, http.StatusCreated)
	Log.HTTPErrorf(request, "test message, args1: %s", "123")
	Log.HTTPWarnf(request, "test message, args1: %s", "123")
	Log.HTTPInfof(request, "test message, args1: %s", "123")

	//Log.PostgresFatalf("test message, args1: %s", "123")
	Log.PostgresErrorf(ctx, "test message, args1: %s", "123")
	Log.PostgresWarnf(ctx, "test message, args1: %s", "123")
	Log.PostgresInfof(ctx, "test message, args1: %s", "123")

	//Log.MongoFatalf("test message, args1: %s", "123")
	Log.MongoErrorf(ctx, "test message, args1: %s", "123")
	Log.MongoWarnf(ctx, "test message, args1: %s", "123")
	Log.MongoInfof(ctx, "test message, args1: %s", "123")

	//Log.RedisFatalf("test message, args1: %s", "123")
	Log.RedisErrorf(ctx, "test message, args1: %s", "123")
	Log.RedisWarnf(ctx, "test message, args1: %s", "123")
	Log.RedisInfof(ctx, "test message, args1: %s", "123")
}

func (suite *LoggerTestSuite) TestRuntimeProcCaptureFailure() {
	expectedFields := logrus.Fields{"file": "unknown", "function": "unknown"}
	actualFields := Log.getProcessFields(uintptr(1), "", 0, false)
	suite.Equal(expectedFields, actualFields)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}
