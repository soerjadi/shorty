package config

import (
	"io/ioutil"

	"github.com/soerjadi/short/internal/pkg/util"
	"gopkg.in/gcfg.v1"
)

var configFilePaths = map[string]string{
	"PRODUCTION":  "/etc/shorty/config.ini",
	"DEVELOPMENT": "../../files/config.ini",
}

func Init() (*Config, error) {
	cfg = &Config{}

	configFilePath := configFilePaths[util.GetENV()]

	config, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = gcfg.ReadStringInto(cfg, string(config))
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// GetConfig returns config object
func GetConfig() *Config {
	return cfg
}
