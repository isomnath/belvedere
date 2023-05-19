package instrumentation

import (
	"log"
	"net"
	"strconv"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/isomnath/belvedere/config"
)

func StartDDTracer() {
	log.Printf("starting data-dog tracer for application: %s", config.GetAppName())
	ddHost := config.GetDataDogTracerConfig().Host()
	if ddHost == "" {
		ddHost = "localhost"
	}
	ddPort := config.GetDataDogTracerConfig().Port()
	if ddPort == 0 {
		ddPort = 8126
	}

	tracer.Start(
		tracer.WithDebugMode(config.GetDataDogTracerConfig().LogLevel() == "DEBUG"),
		tracer.WithAnalytics(true),
		tracer.WithService(config.GetAppName()),
		tracer.WithLogger(datadogLogger{}),
		tracer.WithLogStartup(config.GetDataDogTracerConfig().LogInjectionEnabled()),
		tracer.WithAgentAddr(net.JoinHostPort(ddHost, strconv.Itoa(ddPort))),
	)
}

func StopDDTracer() {
	log.Printf("stopping data-dog tracer for application: %s", config.GetAppName())
	tracer.Stop()
}

type datadogLogger struct{}

func (datadogLogger) Log(str string) {
	log.Printf(str)
}
