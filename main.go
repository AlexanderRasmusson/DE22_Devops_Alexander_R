package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AlexanderRasmusson/DE22_Devops_Alexander_R/server"
)

func main() {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "8080"
	}

	srv := server.New(server.Options{
		Router: http.NewServeMux(),
		Log:    log.Default(),
		Port:   port,
		Host:   os.Getenv("SERVER_HOST"),
	})

	srv.Start()
}
