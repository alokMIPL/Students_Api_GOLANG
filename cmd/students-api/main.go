package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alokMIPL/students-api/internal/config"
)

func main() {

	fmt.Println("Welconme to students api")

	// Load config

	cfg := config.MustLoad()

	log.Println("Environment", cfg.Env)

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welocome to students api"))
	})

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server started %s", cfg.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
