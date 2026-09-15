package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alokMIPL/students-api/internal/config"
)

func main() {

	fmt.Println("Welcome to students api")

	// Load config
	cfg := config.MustLoad()
	log.Println("Environment:", cfg.Env)

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server stared as = ", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
