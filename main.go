package main

import (
	"fmt"

	"github.com/raiki02/video-extractor/internal/appconfig"
	"github.com/raiki02/video-extractor/internal/server"
)

const configPath = "config.yaml"

func main() {
	cfg, err := appconfig.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("load config %s failed: %w", configPath, err))
	}

	e := server.New(cfg)
	if err := e.Run(cfg.Server.Addr); err != nil {
		panic(err)
	}
}
