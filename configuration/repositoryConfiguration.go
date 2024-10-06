package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sync"
)

type Configuration struct{}

var (
	once      sync.Once
	loadError error
)

func LoadConfiguration() error {
	once.Do(func() {
		dir, err := os.UserHomeDir()
		if err != nil {
			loadError = fmt.Errorf("cannot get user home directory: %w", err)
			return
		}
		configPath := path.Join(dir, ".to_remove")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath(configPath)
		err = viper.ReadInConfig()
		if err != nil {
			loadError = fmt.Errorf("cannot read config file: %w", err)
			return
		}
		return
	})
	return loadError
}

func GetString(key string) string {
	return viper.GetString(key)
}
