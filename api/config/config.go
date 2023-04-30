package config

import (
	"log"
	"os"
	"path/filepath"
)

var config *Config

type Config struct {
	AppDir string
}

func GetConfig() Config {
	if config == nil {
		log.Fatal("config was not initialized")
	}
	return *config
}

func LoadConfig() *Config {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	config = &Config{AppDir: filepath.Dir(ex)}
	return config
}
