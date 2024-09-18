package main

import (
	"github.com/UrielVentura/gopportunities/config"
	"github.com/UrielVentura/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
