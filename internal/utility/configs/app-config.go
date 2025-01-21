package configs

import "net/http"


type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ApplicationConfig struct {
	Server Config `json:"server"`
}

func (app *ApplicationConfig) RunApp() error {
	server := http.Server{
		Addr:    app.Server.Host + ":" + app.Server.Port,
	}
	return server.ListenAndServe()
}