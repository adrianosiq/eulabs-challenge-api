package config

import (
	"os"
)

var (
	Cfg *Config
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	DBCharset   string
	DBParseTime string
	DBLoc       string
	PORT        string
}

func LoadConfig() *Config {
	config := &Config{
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_NAME"),
		DBCharset:   os.Getenv("DB_CHARSET"),
		DBParseTime: os.Getenv("DB_PARSETIME"),
		DBLoc:       os.Getenv("DB_LOC"),
		PORT:        os.Getenv("PORT"),
	}

	Cfg = config

	return config
}
