package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Address string   `json:"address"`
	Secret  string   `json:"secret"`
	Nodes   []string `json:"nodes"`
}

var ErrSecretNotChanged = errors.New("please change default secret before use")

func Load() (*Config, error) {
	var cfg Config
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&cfg)
	if cfg.Secret == "changeme" {
		return nil, ErrSecretNotChanged
	}
	return &cfg, err
}

func Default() *Config {
	return &Config{
		Address: "0.0.0.0:8080",
		Secret:  "changeme",
		Nodes:   []string{},
	}
}

func Save(cfg *Config) error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(cfg)
}
