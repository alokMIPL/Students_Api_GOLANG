package main

import (
	"fmt"
	"log"

	"github.com/alokMIPL/students-api/internal/config"
)

func main() {

	// Load config

	cfg := config.MustLoad()

	log.Println("Environment", cfg.Env)

	// database setup
	// setup router
	// setup server

	fmt.Println("Welconme to students api")

}
