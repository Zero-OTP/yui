package config

import (
	"encoding/json"
	"os"
)

// Config es una estructura que representa el archivo de configuración
type Config struct {
	API_URL string `json:"API_URL"`
	DB_NAME string `json:"DB_NAME"`
	DB_USER string `json:"DB_USER"`
	DB_PASS string `json:"DB_PASS"`
}

// cfg es una variable global que almacena la configuración cargada
var cfg *Config

// Load carga el archivo de configuración desde el disco
func Load() error {
	file, err := os.Open("./config/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	cfg = &Config{}
	if err := json.NewDecoder(file).Decode(cfg); err != nil {
		return err
	}

	return nil
}

// GetConfig devuelve la configuración cargada
func GetConfig() *Config {
	return cfg
}
