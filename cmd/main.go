package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "github.com/Ademayowa/kubernetes-job-board/internal/database"
	"github.com/Ademayowa/kubernetes-job-board/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.InitDB()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	handlers.RegisterRoutes(router)

	srv := &http.Server{Addr: ":" + port, Handler: router}

	// Start server in background
	go srv.ListenAndServe()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Server stopped")
}
