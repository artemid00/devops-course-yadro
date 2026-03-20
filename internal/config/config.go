package config

import (
	"os"
	"unicode"
)

type Config struct {
	Port    string
	Version string
	Author  string
}

func isDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return s != ""
}

// loads parameters from env
func LoadConfig() Config {
	var cfg Config

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" || !isDigits(cfg.Port) {
		cfg.Port = "8000"
	}

	cfg.Version = os.Getenv("VERSION")
	if cfg.Version == "" {
		cfg.Version = "1.1.0"
	}

	cfg.Author = os.Getenv("AUTHOR")
	if cfg.Author == "" {
		cfg.Author = "a.bezpyatko"
	}

	return cfg
}
