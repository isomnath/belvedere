package config

type SentryConfig struct {
	enabled bool
	dsn     string
}

func sentryConfig() *SentryConfig {
	return &SentryConfig{
		enabled: getBool(sentryEnabled, false),
		dsn:     getString(sentryDSN, false),
	}
}

func (sentry *SentryConfig) Enabled() bool {
	return sentry.enabled
}

func (sentry *SentryConfig) DSN() string {
	return sentry.dsn
}
