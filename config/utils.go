package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func getInt(key string, required bool) int {
	if required {
		checkKey(key)
	}
	return viper.GetInt(key)
}

func getString(key string, required bool) string {
	if required {
		checkKey(key)
	}
	return viper.GetString(key)
}

func getBool(key string, required bool) bool {
	if required {
		checkKey(key)
	}
	return viper.GetBool(key)
}

func getStringSlice(key string, required bool) []string {
	if required {
		checkKey(key)
	}

	envVal := strings.ReplaceAll(viper.GetString(key), " ", "")
	return strings.Split(envVal, ",")
}

func getIntSlice(key string, required bool) []int {
	if required {
		checkKey(key)
	}

	os.Getenv(key)
	envVal := strings.ReplaceAll(viper.GetString(key), " ", "")
	stringVal := strings.Split(envVal, ",")

	var intSlice []int
	for _, str := range stringVal {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			log.Panicf("Cannot convert value to integer array for key: %s", key)
		}
		intSlice = append(intSlice, intVal)
	}
	return intSlice
}

func getStringMapInterface(key string, required bool) map[string]interface{} {
	if required {
		checkKey(key)
	}
	return viper.GetStringMap(key)
}

func getStringMapString(key string, required bool) map[string]string {
	if required {
		checkKey(key)
	}
	return viper.GetStringMapString(key)
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		log.Panicf("Missing key: %s", key)
	}
}
