package config

import (
	"os"
)

// Config ...
type Config struct {
	AppPort          string
	DBHost           string
	DBUsername       string
	DBPassword       string
	DBName           string
	DBPort           string
	SessionAPIHost   string
	SalesTeamAPIHost string
	DBMedeaHost      string
	DBMedeaUsername  string
	DBMedeaPassword  string
	DBMedeaName      string
	DBMedeaPort      string
}

// NewConfig ...
func NewConfig() Config {
	config := Config{
		AppPort:          os.Getenv("APP_PORT"),
		DBHost:           os.Getenv("DB_HOST"),
		DBUsername:       os.Getenv("DB_USERNAME"),
		DBPassword:       os.Getenv("DB_PASSWORD"),
		DBName:           os.Getenv("DB_NAME"),
		DBPort:           os.Getenv("DB_PORT"),
		SessionAPIHost:   os.Getenv("SESSION_API_HOST"),
		SalesTeamAPIHost: os.Getenv("SALES_TEAM_API_HOST"),
		DBMedeaHost:      os.Getenv("DB_MEDEA_HOST"),
		DBMedeaUsername:  os.Getenv("DB_MEDEA_USERNAME"),
		DBMedeaPassword:  os.Getenv("DB_MEDEA_PASSWORD"),
		DBMedeaName:      os.Getenv("DB_MEDEA_NAME"),
		DBMedeaPort:      os.Getenv("DB_MEDEA_PORT"),
	}

	if config.AppPort == "" {
		config.AppPort = "8080"
	}

	return config
}
