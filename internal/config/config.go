package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	DBBM   DBConfig
	DBSdk  DBConfig
	DBMkt  DBConfig
	JWT    JWTConfig
	Mongo  MongoConfig
	Redis  RedisConfig
	OAuth  OAuthConfig
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

type MongoConfig struct {
	URI        string
	ChannelDB  string
	ChannelRaw string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type OAuthConfig struct {
	RedirectPath string
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
		DBSdk: DBConfig{
			Host:     viper.GetString("DB_SDK_HOST"),
			Port:     viper.GetString("DB_SDK_PORT"),
			User:     viper.GetString("DB_SDK_USER"),
			Password: viper.GetString("DB_SDK_PASSWORD"),
			Name:     viper.GetString("DB_SDK_NAME"),
		},
		DBMkt: DBConfig{
			Host:     viper.GetString("DB_MKT_HOST"),
			Port:     viper.GetString("DB_MKT_PORT"),
			User:     viper.GetString("DB_MKT_USER"),
			Password: viper.GetString("DB_MKT_PASSWORD"),
			Name:     viper.GetString("DB_MKT_NAME"),
		},
		JWT: JWTConfig{
			Secret:      viper.GetString("JWT_SECRET"),
			ExpireHours: viper.GetInt("JWT_EXPIRE_HOURS"),
		},
		Mongo: MongoConfig{
			URI:        viper.GetString("MONGO_URI"),
			ChannelDB:  viper.GetString("MONGO_CHANNEL_DB"),
			ChannelRaw: viper.GetString("MONGO_CHANNEL_RAW_DB"),
		},
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		OAuth: OAuthConfig{
			RedirectPath: viper.GetString("OAUTH_REDIRECT_PATH"),
		},
	}

	return AppConfig
}

func (d DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.Name)
}
