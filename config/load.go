package config

import (
	"fmt"
	"github.com/joho/godotenv"
	log2 "github.com/labstack/gommon/log"
	"log"
	"os"
	"strings"
)

type Config struct {
	HankoApiUrl  string
	DatabaseName string
	LogLevel     log2.Lvl
}

// LoadConfig
// Loads the config and creates a Config struct
// It will load some values from env variables
// And some can be loaded from config files
func LoadConfig() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	hankoApiUrl, ok := os.LookupEnv("HANKO_API_URL")
	if !ok || hankoApiUrl == "" {
		log.Fatal("'HANKO_API_KEY' is not set")
	}

	databaseName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok || databaseName == "" {
		log.Fatal("'DATABASE_NAME' is not set")
	}

	var logLvl log2.Lvl
	logLvlEnv, ok := os.LookupEnv("LOG_LEVEL")
	if !ok || logLvlEnv == "" {
		log.Println("WARNING: log level not set using default log2.WARN")
		logLvl = log2.WARN
	} else {
		logLvl, err = stringToLevel(logLvlEnv)
		if err != nil {
			log.Fatal(err)
		}
	}

	return Config{
		HankoApiUrl:  hankoApiUrl,
		DatabaseName: databaseName,
		LogLevel:     logLvl,
	}
}

func stringToLevel(level string) (log2.Lvl, error) {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return log2.DEBUG, nil
	case "INFO":
		return log2.INFO, nil
	case "WARN":
		return log2.WARN, nil
	case "ERROR":
		return log2.ERROR, nil
	case "OFF":
		return log2.OFF, nil
	default:
		return 0, fmt.Errorf("invalid log level: %s", level)
	}
}
