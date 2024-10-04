package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort    string
	Dsn           string
	Secret        string
	ElasticUrl    string
	PrometheusUrl string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		log.Println("HTTP_PORT not found using default port :3000")
		httpPort = ":3000"
	}

	Dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER_NAME"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	fmt.Println(Dsn)
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("DSN variables not found")
	}

	appSecret := os.Getenv("SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("app secret not found")
	}

	elasticUrl := os.Getenv("ELASTICSEARCH_URL")
	if len(elasticUrl) < 1 {
		return AppConfig{}, errors.New("elastic url not found")
	}

	prometheusUrl := os.Getenv("PROMETHEUS_URL")
	if len(prometheusUrl) < 1 {
		return AppConfig{}, errors.New("prometheus url not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, Secret: appSecret, ElasticUrl: elasticUrl, PrometheusUrl: prometheusUrl}, nil
}
