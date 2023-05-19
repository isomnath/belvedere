package config

type NewRelicConfig struct {
	enabled    bool
	licenseKey string
}

func newRelicConfig() *NewRelicConfig {
	return &NewRelicConfig{
		enabled:    getBool(newRelicEnabled, false),
		licenseKey: getString(newRelicLicenseKey, false),
	}
}

func (nr *NewRelicConfig) Enabled() bool {
	return nr.enabled
}

func (nr *NewRelicConfig) LicenseKey() string {
	return nr.licenseKey
}
