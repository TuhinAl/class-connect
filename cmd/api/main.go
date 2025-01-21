package main

import (
	"golang-api/cmd/api/handler"
	"golang-api/internal/utility/configs"
	"log"
	"net/http"
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
	
	log.Fatal(applicationConfig.RunApp(externalRoutes()))
}


func  externalRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.ApplicationHealthHandler)
	mux.HandleFunc("GET /student", handler.GetStudentHandler)
	return mux
}