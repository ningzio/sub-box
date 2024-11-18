package main

import (
	"log"
	"net/http"

	"github.com/ningzio/sub-box/backend/internal/api"
	"github.com/ningzio/sub-box/backend/internal/config"
)

func main() {
	cfg := config.Load()

	server := api.NewServer(cfg)

	log.Printf("Server starting on %s", cfg.ListenAddr)
	if err := http.ListenAndServe(cfg.ListenAddr, server.Router()); err != nil {
		log.Fatal(err)
	}
}
