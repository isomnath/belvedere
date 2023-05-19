package config

type TranslationConfig struct {
	enabled            bool
	path               string
	whitelistedLocales []string
	defaultLocale      string
}

func translationsConfig() *TranslationConfig {
	return &TranslationConfig{
		enabled:            getBool(translationsEnabled, false),
		path:               getString(translationsPath, false),
		whitelistedLocales: getStringSlice(translationsWhitelistedLocales, false),
		defaultLocale:      getString(translationsDefaultLocale, false),
	}
}

func (tr *TranslationConfig) Enabled() bool {
	return tr.enabled
}

func (tr *TranslationConfig) Path() string {
	return tr.path
}
func (tr *TranslationConfig) WhitelistedLocales() []string {
	return tr.whitelistedLocales
}
func (tr *TranslationConfig) DefaultLocale() string {
	return tr.defaultLocale
}
