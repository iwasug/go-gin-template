package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

// GetConfig By Env
func GetConfig() *viper.Viper {
	return config
}

var env *viper.Viper

// GetEnv App
func GetEnv() string {
	var err error
	env = viper.New()
	env.SetConfigType("yaml")
	env.SetConfigName("env")
	env.AddConfigPath("../config/")
	env.AddConfigPath("config/")
	err = env.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	envApp := env.GetString("environment")
	return envApp
}
