package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

// FOR LOAD ALL CONFIG

type Config struct {
	Server
	DataBase
	Elastic
}

type Server struct {
	Host string
	Port string
}

type DataBase struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type Elastic struct {
	Host         string
	Username     string
	Password     string
	MaxIdleConns int
	Timeout      time.Duration
}

// <-- CONSTRUCTOR --> //

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in load .env : ", err.Error())
	}

	// Parse ELASTICSEARCH_MAX_IDLE_CONNS as integer
	maxIdleConns, err := strconv.Atoi(os.Getenv("ELASTICSEARCH_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatalf("Error converting ELASTICSEARCH_MAX_IDLE_CONNS to int: %v", err)
	}

	// Parse ELASTICSEARCH_TIMEOUT as duration
	timeout, err := time.ParseDuration(os.Getenv("ELASTICSEARCH_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error parsing ELASTICSEARCH_TIMEOUT: %v", err)
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DataBase: DataBase{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
		Elastic: Elastic{
			Host:         os.Getenv("ELASTICSEARCH_HOST"),
			Username:     os.Getenv("ELASTICSEARCH_USERNAME"),
			Password:     os.Getenv("ELASTICSEARCH_PASSWORD"),
			MaxIdleConns: maxIdleConns,
			Timeout:      timeout,
		},
	}

}
