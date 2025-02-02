package main

import (
	"golang-api/cmd/api/handler"
	dbconfig "golang-api/internal/db-config"
	"golang-api/internal/repository"
	
	"log"
	_ "github.com/lib/pq"
)

// import "golang-api/internal/utility/configs"

func main() {
	app := handler.ApplicationConfig{
		Server: handler.Config{
			Host: "localhost",
			Port: "8080",
			DB: dbconfig.DBConfig{
				Addr:         "postgres://postgres:postgres@localhost:5432/class_connect?sslmode=disable",
				MaxOpenConns: 20,
				MaxIdleConns: 20,
				MaxIdleTime:  "15m",
			},
			SMTP: handler.SMTP{
				Host:     "smtp.gmail.com",
				Port:     465,
				Username: "a033e9fca16213",
				Password: "ad86d311e28d8b",
				Sender:   "your_sender_email",
			},
		},
	}

	db, err := dbconfig.NewDBConfig(app.Server.DB.Addr,
		app.Server.DB.MaxOpenConns,
		app.Server.DB.MaxIdleConns,
		app.Server.DB.MaxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connected successfully")

	app.Store = repository.NewStorage(db)

	log.Fatal(app.RunApp(handler.ExternalRoutes(&app)))
}

// func externalRoutes() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET /health", handler.ApplicationHealthHandler)
// 	mux.HandleFunc("GET /student", handler.GetStudentHandler)
// 	return mux
// }

// func externalRoutes() http.Handler {
// 	router := chi.NewRouter()

// 	// A good base middleware stack
// 	router.Use(middleware.RequestID)
// 	router.Use(middleware.RealIP)
// 	router.Use(middleware.Logger)
// 	router.Use(middleware.Recoverer)

// 	// Set a timeout value on the request context (ctx), that will signal
// 	// through ctx.Done() that the request has timed out and further
// 	// processing should be stopped.
// 	router.Use(middleware.Timeout(60 * time.Second))

// 	router.Route("/v1", func(router chi.Router) {
// 		router.Get("/health", handler.ApplicationHealthHandler)
// 		router.Route("/student", func(router chi.Router) {
// 			router.Get("/", handler.GetStudentHandler)
// 			router.Post("/create", handler.CreateStudentHandler)
// 		})
// 	})

// 	return router
// }
