package config

import "time"

type DataDogConfig struct {
	enabled            bool
	host               string
	port               int
	flushPeriodSeconds time.Duration
}

func dataDogConfig() *DataDogConfig {
	return &DataDogConfig{
		enabled:            getBool(dataDogEnabled, false),
		host:               getString(dataDogHost, false),
		port:               getInt(dataDogPort, false),
		flushPeriodSeconds: time.Duration(getInt(dataDogFlushPeriodSeconds, false)),
	}
}

func (dataDog *DataDogConfig) Enabled() bool {
	return dataDog.enabled
}

func (dataDog *DataDogConfig) Host() string {
	return dataDog.host
}

func (dataDog *DataDogConfig) Port() int {
	return dataDog.port
}

func (dataDog *DataDogConfig) FlushPeriod() time.Duration {
	return dataDog.flushPeriodSeconds
}
