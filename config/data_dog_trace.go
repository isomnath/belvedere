package config

type DataDogTraceConfig struct {
	enabled             bool
	host                string
	port                int
	logLevel            string
	logInjectionEnabled bool
}

func dataDogTraceConfig() *DataDogTraceConfig {
	return &DataDogTraceConfig{
		enabled:             getBool(dataDogTraceAgentEnabled, false),
		host:                getString(dataDogTraceAgentHost, false),
		port:                getInt(dataDogTraceAgentPort, false),
		logLevel:            getString(dataDogTraceLogLevel, false),
		logInjectionEnabled: getBool(dataDogLogInjectionEnabled, false),
	}
}

func (dataDogTracer *DataDogTraceConfig) Enabled() bool {
	return dataDogTracer.enabled
}

func (dataDogTracer *DataDogTraceConfig) Host() string {
	return dataDogTracer.host
}

func (dataDogTracer *DataDogTraceConfig) Port() int {
	return dataDogTracer.port
}

func (dataDogTracer *DataDogTraceConfig) LogLevel() string {
	return dataDogTracer.logLevel
}

func (dataDogTracer *DataDogTraceConfig) LogInjectionEnabled() bool {
	return dataDogTracer.logInjectionEnabled
}
