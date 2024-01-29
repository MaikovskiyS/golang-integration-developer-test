package main

import (
	"log"

	"integration.v1/internal/app"
	"integration.v1/internal/config"
)

func main() {
	cfg := config.New()
	err := app.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
