package main

import (
	"log"

	"github.com/millirud/go-scrapper/config"
	"github.com/millirud/go-scrapper/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
