package main

import (
	"github.com/devlucas-java/api-opportunities/config"
	r "github.com/devlucas-java/api-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("API-OPPORTUNITIES")

	err := config.Init()
	if err != nil {
		logger.Debugf("Error initializing config: %v", err)
		return
	}

	// start all over
	r.Initalizer()
}
