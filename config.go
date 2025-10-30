package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func Load[T any](configFile string, defaultConfig ...T) T {
	var config T
	if len(defaultConfig) >= 1 {
		config = defaultConfig[0]
	}

	_ = godotenv.Load(".env")
	_ = loadConfigFile(&config, configFile)
	_ = loadConfigEnv(&config)

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
