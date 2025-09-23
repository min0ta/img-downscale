package main

import (
	"flag"
	"fmt"
	"image_thumb/internal/api"
	"image_thumb/internal/config"
	"os"
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.json", "path to json config file")
	flag.StringVar(&configPath, "c", "config/config.json", "path to json config file")
}

var (
	configPath string
)

func main() {
	flag.Parse()
	cfg := config.GetDefault()
	cfg.ReadJsonConfig(configPath)

	server, err := api.New(cfg)
	if err != nil {
		fmt.Println("unable create server with error %v", err)
		os.Exit(1)
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("unable listen and serve with error %v", err)
		os.Exit(1)
	}
}
