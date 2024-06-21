package main

import (
	"github.com/JMustang/jobsopportunities/config"
	"github.com/JMustang/jobsopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Inicializando Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
