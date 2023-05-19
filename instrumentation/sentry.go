package instrumentation

import (
	"log"

	"github.com/getsentry/sentry-go"

	"github.com/isomnath/belvedere/config"
)

var client *sentry.Client

func InitializeSentry(config *config.SentryConfig) {
	if config.Enabled() {
		var err error
		options := sentry.ClientOptions{
			Dsn: config.DSN(),
		}
		client, err = sentry.NewClient(options)

		if err != nil {
			log.Panicf("failed initializing sentry client with error: %v", err)
		}
	}
}

func GetSentryClient() *sentry.Client {
	return client
}

func CaptureError(err error) {
	if client != nil {
		client.CaptureException(err, nil, nil)
	}
}

func CaptureErrorWithTags(err error, tags map[string]string) {
	scope := sentry.NewScope()
	scope.SetTags(tags)
	if client != nil {
		client.CaptureException(err, nil, scope)
	}
}

func CaptureWarn(err error) {
	scope := sentry.NewScope()
	scope.SetTags(map[string]string{
		"level": string(sentry.LevelWarning),
	})
	if client != nil {
		client.CaptureException(err, nil, nil)
	}
}
