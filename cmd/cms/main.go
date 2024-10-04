package main

import (
	"github.com/hritesh04/news-system/config"
	"github.com/hritesh04/news-system/internal/api"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		panic(err)
	}
	api.StartServer(cfg)
}
