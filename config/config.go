package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/mgorunuch/environment-launcher/bindata"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

const CONFIG_FILE_PERM = 0764

type Command struct {
	Name  string       `yaml:"name"`
	Shell ShellCommand `yaml:"shell"`
}

type App struct {
	Name     string    `yaml:"name"`
	Commands []Command `yaml:"commands"`
}

type Config struct {
	Apps       []App `yaml:"apps"`
	ConfigPath string
}

func InitConfig(logger *zap.Logger, filePath string) (config *Config, err error) {
	var (
		contents []byte
	)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user $HOME")
	}

	filePath = path.Join(homeDir, ".config", "environment-launcher", "config.yaml")
	contents, err = getFileOrCreateNew(filePath)

	config = &Config{
		ConfigPath: filePath,
	}
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		logger.Error("failed to unmarshal file", zap.Error(err))
		return nil, err
	}

	return config, nil
}

// Do not log error
func getFileOrCreateNew(path string) (contents []byte, err error) {
	contents, err = ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		dir := filepath.Dir(path)

		err = os.MkdirAll(dir, CONFIG_FILE_PERM)
		if err != nil {
			return nil, err
		}

		contents, err = bindata.Asset("static/config.yaml")
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(path, contents, CONFIG_FILE_PERM)
		if err != nil {
			return nil, err
		}
	}

	return contents, nil
}
