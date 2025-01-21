package configs

import (
	"log"
	"net/http"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ApplicationConfig struct {
	Server Config `json:"server"`
}

func (app *ApplicationConfig) RunApp() error {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    app.Server.Host + ":" + app.Server.Port,
		Handler: mux,
	}
	

	log.Printf("Server starting on http://%s:%s", app.Server.Host, app.Server.Port)
	log.Printf("====SERVER STARTED SUCCESSFULLY===")
	return server.ListenAndServe()
}
