package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// TODO: Add tests to handle error scenarios
func customConfig(val interface{}) {
	envKeysMap := &map[string]interface{}{}
	if err := mapstructure.Decode(val, &envKeysMap); err != nil {
		log.Panicf("failed to decode environment keys to struct: %v", err)
	}

	for k := range *envKeysMap {
		if bindErr := viper.BindEnv(k); bindErr != nil {
			log.Panicf("failed to bind environment keys: %v", bindErr)
		}
	}

	if unmarshalErr := viper.Unmarshal(val); unmarshalErr != nil {
		log.Panicf("custom Config unmarshal failed for struct: %v with error: %v", val, unmarshalErr)
	}
}
