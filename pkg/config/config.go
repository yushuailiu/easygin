package config

import (
	"os"
	"path/filepath"
	"strings"
	goconfig "github.com/yushuailiu/easygin/config"
	"io/ioutil"
	"github.com/go-ini/ini"
)

var cfg *ini.File

type Config struct {
	BasePath	string
	ConfigPath	string
}

func DefaultConfig() *Config {
	return &Config{
		BasePath: "/config",
	}
}

func GetConfig() *ini.File {
	return cfg
}

func (config *Config) Bootstrap(env string) *ini.File {
	configPath := config.BasePath + "/" + env
	config.ConfigPath = configPath

	cfg = ini.Empty()

	for name, file := range goconfig.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".ini") || !strings.HasPrefix(name, configPath) {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			panic("load config file fail:" + name)
		}
		cfg.Append(h)
	}

	return cfg
}

func (config *Config) getAllIniFiles() []string {
	path := config.ConfigPath
	paths := make([]string, 0)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".ini") {
			paths = append(paths, path)
			return nil
		}
		return err
	})
	return paths
}
