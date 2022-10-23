package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env              string `mapstructure:"env_prefix"`
	ServerPort       int    `mapstructure:"server_port"`
	DBConnection     string `mapstructure:"db_connection"`
	DefaultPageNum   int    `mapstructure:"default_page_num"`
	DefaultPageLimit int    `mapstructure:"default_page_limit"`
	LogConfig        LogConfig
}

type LogConfig struct {
	Format  string `mapstructure:"format"`
	Output  string `mapstructure:"output"`
	Expired int    `mapstructure:"expired"`
	Path    string `mapstructure:"path"`
}

//InitFromFile init config file
func InitFromFile(path string) *Config {
	cfg := new(Config)
	if path == "" {
		viper.AddConfigPath("config")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	} else {
		viper.SetConfigFile(path)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Config file not found: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("covert to struct: %v", err)
	}

	if path == "" {
		fmt.Printf("File config used  %s\n \n", viper.ConfigFileUsed())
		dataPrinf, _ := json.Marshal(cfg)
		fmt.Printf("Config:  %s\n \n", string(dataPrinf))
	}

	return cfg
}
