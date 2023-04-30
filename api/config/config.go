package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var config *Config

type Config struct {
	AppDir        string `toml:"-"`
	ConsoleFolder string `toml:"console_folder"`
}

func GetConfig() Config {
	if config == nil {
		LoadConfig()
	}
	return *config
}

func LoadConfig() *Config {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	config = &Config{AppDir: filepath.Dir(ex)}
	if _, err := os.Stat(filepath.Join(config.AppDir, "pfile.toml")); err != nil {

		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		config.AppDir = filepath.Join(path, "dist")
	}
	if _, err := toml.DecodeFile(filepath.Join(config.AppDir, "pfile.toml"), config); err != nil {
		log.Fatal(err)
	}

	return config
}