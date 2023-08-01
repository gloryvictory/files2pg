package config

import "github.com/spf13/viper"

type DbConfig struct {
	Address  string `mapstructure:"address"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"user"`
	Schema   string `mapstructure:"schema"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}
type SourceConfig struct {
	Folder string `mapstructure:"folder"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
	Source SourceConfig `mapstructure:"source"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config
	vp.SetConfigName("files2pg")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("..")
	vp.AddConfigPath("./internal/config")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
