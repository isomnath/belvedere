package translator

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type translator struct {
	bundle *i18n.Bundle
}

var t *translator

// Initialize - Initializes the translation bundle
func Initialize() {
	t = &translator{bundle: loadTranslations()}
}

func Kill() {
	t = nil
}

func loadTranslations() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	translationsPath := config.GetTranslationConfig().Path()

	err := filepath.Walk(translationsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(info.Name(), ".json") {
			_, err := bundle.LoadMessageFile(path)
			if err != nil {
				log.Log.Errorf(context.Background(), "error while loading translation file: %v", err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Log.Panicf("failed to load translations: %v", err)
	}

	return bundle
}

// Translate - Translates message to supported language bundle
func Translate(message, language string) string {
	if t == nil {
		return message
	}

	l := i18n.NewLocalizer(t.bundle, language)

	translatedMessage, err := l.Localize(&i18n.LocalizeConfig{MessageID: message})

	if err != nil {
		log.Log.Warnf(context.Background(), "failed while translating message with ID: %s with error: %v", message, err)
		return message
	}
	return translatedMessage
}
