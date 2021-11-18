package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/workfoxes/calypso/pkg/log"
)

var C *Config

var config map[string]interface{}

type Config struct {
	TraderKey          string `json:"TraderKey"`
	Env                string `json:"ENV"`
	GoogleClientId     string `json:"GoogleClientId"`
	GoogleClientSecret string `json:"GoogleClientSecret"`
	HOST               string `json:"Host"`
	RedisHost          string `json:"RedisHost"`
	RedisPassword      string `json:"RedisPassword"`
}

var DefaultConfigFileName = "config.json"

func GetConfig() *Config {
	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "Dev"
	}
	var _config Config
	data, _ := ioutil.ReadFile(DefaultConfigFileName)
	_ = json.Unmarshal(data, &config)
	_ = json.Unmarshal(data, &_config)
	config["Env"] = env
	return &_config
}

func Get(key string) interface{} {
	return config[key]
}

func GetDefault(key string, _default interface{}) interface{} {
	_value := Get(key)
	if _value == nil {
		return _default
	}
	return _value
}

func GetBool(key string) bool {
	return Get(key).(bool)
}

func GetEnv(key string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		log.Panic("ENV not found %s", key)
		return ""
	}
	return env
}
