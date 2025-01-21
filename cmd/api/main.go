package main

import (
	"golang-api/internal/utility/configs"
	"log"
)

// import "golang-api/internal/utility/configs"

func main() {
	// configs.NewAPIServer(":8080").Run()

	applicationConfig := &configs.ApplicationConfig{
		Server: configs.Config{
			Host: "localhost",
			Port: "8080",
		},
	}

	log.Fatal(applicationConfig.RunApp())
}
