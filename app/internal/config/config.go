package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Config struct {
	Port            string
	StoragePath     string
	MigrationsPath  string
	DBConnectionURL string
	DownloadLimit   int64
	MaxRoutines     int
}

func GetDefault() *Config {
	return &Config{
		Port:            ":8080",
		DBConnectionURL: "postgresql://postgres:root@localhost:5432/thumb",
		StoragePath:     "./storage/",
		MigrationsPath:  "./migrations/migrations.sql",
		DownloadLimit:   30 * 1024 * 1024,
		MaxRoutines:     5,
	}
}

func (cfg *Config) ReadJsonConfig(path string) {
	configFile, err := os.ReadFile(path)
	if err == nil {
		cfg.assignJsonValue(configFile)
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		cfg.createDefaultConfigFile()
	} else {
		fatalifnil("cannot open config file with error", err)
	}
}

func (cfg *Config) assignJsonValue(configFile []byte) {
	err := json.Unmarshal(configFile, &cfg)
	if err != nil {
		fatalifnil("couldnt parse json from config file with error", err)
	}
}

func (cfg *Config) createDefaultConfigFile() {
	err := os.MkdirAll("config", os.ModePerm)
	fatalifnil("cannot create config dir with error", err)

	file, err := os.Create("config/config.json")
	fatalifnil("cannot create config file", err)
	defer file.Close()

	jsonConfig, err := json.Marshal(cfg)
	fatalifnil("cannot marshal cfg ", err)

	file.Write(jsonConfig)
	configFile := jsonConfig

	fatalifnil("unable to serialize json config with error: ", json.Unmarshal(configFile, cfg))
}

func fatalifnil(explanation string, err error) {
	if err != nil {
		log.Fatal(explanation, err)
	}
}
