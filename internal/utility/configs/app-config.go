package configs

import (
	dbconfig "golang-api/internal/db-config"
	"golang-api/internal/repository"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   dbconfig.DBConfig `json:"db"`
}

type ApplicationConfig struct {
	Server Config `json:"server"`
	Store  repository.Storage
	
}

// func (app *ApplicationConfig) RunApp(mux *http.ServeMux) error {
// 	server := http.Server{
// 		Addr:         app.Server.Host + ":" + app.Server.Port,
// 		Handler:      mux,
// 		ReadTimeout:  time.Second * 10,
// 		WriteTimeout: time.Second * 30,
// 		IdleTimeout:  time.Second * 60,
// 	}

// 	log.Printf("Server starting on http://%s:%s", app.Server.Host, app.Server.Port)
// 	log.Printf("====SERVER STARTED SUCCESSFULLY===")
// 	return server.ListenAndServe()
// }

func (app *ApplicationConfig) RunApp(mux http.Handler) error {
	server := http.Server{
		Addr:         app.Server.Host + ":" + app.Server.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server starting on http://%s:%s", app.Server.Host, app.Server.Port)
	log.Printf("====SERVER STARTED SUCCESSFULLY===")
	return server.ListenAndServe()
}
