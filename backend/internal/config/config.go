package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ListenAddr string `json:"listen_addr"`
	Debug      bool   `json:"debug"`
}

func Load() *Config {
	cfg := &Config{
		ListenAddr: ":8080",
		Debug:      true,
	}

	if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
		file, err := os.Open(configPath)
		if err != nil {
			log.Printf("Warning: Could not open config file: %v", err)
			return cfg
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(cfg); err != nil {
			log.Printf("Warning: Could not decode config file: %v", err)
			return cfg
		}
	}

	return cfg
}
