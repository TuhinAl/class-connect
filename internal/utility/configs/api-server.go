package configs

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
	
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		Host: "localhost",
		Port: "8080",
	}
}

func NewAPIServerHostPort(addr string, port string) *APIServer {
	return &APIServer{
		Host: addr,
		Port: port,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	//add handlers
	router.HandleFunc("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("userId")
		w.Write([]byte("User ID: " + userId))

	})


	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router,
	}
	fmt.Println()

	log.Printf("Server is listening on http://%s:%s", s.Host, s.Port)
	return server.ListenAndServe()
}


