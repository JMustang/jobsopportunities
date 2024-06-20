package main

import (
	"fmt"

	"github.com/JMustang/jobsopportunities/config"
	"github.com/JMustang/jobsopportunities/router"
)

func main() {

	// Inicializando Config
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
