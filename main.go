package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"test/database"
	r "test/router"
)

func init() {
	config, err := database.GetDBProperties()
	if err != nil {
		log.Fatal("Cannot read data from config file")
	}
	if ok := database.InitializeDB(config); !ok {
		log.Fatal("Cannot connect db")
	}
}

func main() {
	fmt.Println("Starting server :6000")
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	http.ListenAndServe(":6000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r.Router))
}
