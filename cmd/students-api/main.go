package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alokMIPL/students-api/internal/config"
	"github.com/alokMIPL/students-api/internal/http/handlers/student"
	"github.com/alokMIPL/students-api/internal/http/middleware"
	"github.com/alokMIPL/students-api/internal/storage/sqlite"
)

func main() {

	fmt.Println("Welcome to students api")

	// Load config
	cfg := config.MustLoad()
	log.Println("Environment:", cfg.Env)

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	log.Println("storage initialized", slogAddr(cfg))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	router.HandleFunc("POST /students", student.New(storage))
	router.HandleFunc("GET /students/{id}", student.GetById(storage))
	router.HandleFunc("GET /students", student.GetList(storage))

	// wrap router with middleware chain
	handler := middleware.Chain(router,
		middleware.RecoveryMiddleware,
		middleware.LoggingMiddleware,
		middleware.CORSMiddleware,
	)

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	// Run server in a goroutine so it doesn't block shutdown handling
	go func() {
		fmt.Printf("Server started on %s\n", cfg.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server:", err)
		}
	}()

	// Wait for interrupt signal (Ctrl+C) or termination signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.Println("shutting down server...")

	// Give in-flight requests up to 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown server gracefully:", err)
	}

	log.Println("server stopped gracefully")
}

func slogAddr(cfg *config.Config) string {
	return cfg.StoragePath
}
