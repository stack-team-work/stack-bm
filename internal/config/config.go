package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	DBBM   DBConfig
	DBApi  DBConfig
	JWT    JWTConfig
	Dev    DevConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type DevConfig struct {
	DefaultUsername string
	DefaultPassword string
}

var AppConfig *Config

func LoadConfig() *Config {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: could not read .env: %v\n", err)
	}

	AppConfig = &Config{
		Server: ServerConfig{
			Port: viper.GetString("SERVER_PORT"),
			Mode: viper.GetString("SERVER_MODE"),
		},
		DBBM: DBConfig{
			Host:     viper.GetString("DB_BM_HOST"),
			Port:     viper.GetString("DB_BM_PORT"),
			User:     viper.GetString("DB_BM_USER"),
			Password: viper.GetString("DB_BM_PASSWORD"),
			Name:     viper.GetString("DB_BM_NAME"),
		},
		DBApi: DBConfig{
			Host:     viper.GetString("DB_API_HOST"),
			Port:     viper.GetString("DB_API_PORT"),
			User:     viper.GetString("DB_API_USER"),
			Password: viper.GetString("DB_API_PASSWORD"),
			Name:     viper.GetString("DB_API_NAME"),
		},
		JWT: JWTConfig{
			Secret:      viper.GetString("JWT_SECRET"),
			ExpireHours: viper.GetInt("JWT_EXPIRE_HOURS"),
		},
		Dev: DevConfig{
			DefaultUsername: viper.GetString("DEV_DEFAULT_USERNAME"),
			DefaultPassword: viper.GetString("DEV_DEFAULT_PASSWORD"),
		},
	}

	return AppConfig
}

func (d DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.Name)
}
