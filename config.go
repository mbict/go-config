package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

func Load[T any](configFile string, defaultConfig ...T) T {
	var config T
	if len(defaultConfig) >= 1 {
		config = defaultConfig[0]
	}

	loadConfigFile(&config, configFile)
	loadConfigEnv(&config)

	return config
}

func loadConfigEnv[T any](cfg *T) error {
	return envconfig.Process("", cfg)
}

func loadConfigFile[T any](cfg *T, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	return decoder.Decode(cfg)
}
