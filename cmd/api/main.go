package main

import (
	"golang-api/cmd/api/handler"
	"golang-api/internal/repository"
	"golang-api/internal/utility/configs"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// import "golang-api/internal/utility/configs"

func main() {
	// configs.NewAPIServer(":8080").Run()

	applicationConfig := &configs.ApplicationConfig{
		Server: configs.Config{
			Host: "localhost",
			Port: "8080",
		},
		Store: repository.NewStorage(nil),
	}

	log.Fatal(applicationConfig.RunApp(externalRoutes()))
}

// func externalRoutes() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET /health", handler.ApplicationHealthHandler)
// 	mux.HandleFunc("GET /student", handler.GetStudentHandler)
// 	return mux
// }

func externalRoutes() http.Handler {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/v1", func(router chi.Router) {
		router.Get("/health", handler.ApplicationHealthHandler)
		router.Get("/student", handler.GetStudentHandler)
	})

	return router
}
