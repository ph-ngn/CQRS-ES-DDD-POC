package main

import (
	"context"
	"github.com/andyj29/wannabet/internal/httprest"
	prodLog "github.com/andyj29/wannabet/internal/log"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	server, err := httprest.NewServer(httprest.Config{
		Address: os.Getenv("SERVICE_ADDR"),
		Logger:  prodLog.GetLogger(),
		Router:  chi.NewRouter(),
	})
	if err != nil {
		log.Fatalf("Failed to initialize new HTTP server")
	}

	go func() {
		if err := server.ListenandServe(); err != nil {
			log.Fatalf("Fail to listen and serve %v", err)
		}
	}()

	onSystemSignal(server)
}

func onSystemSignal(server *httprest.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigs:
		log.Printf("System signal detected %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to gracefully shut down server: ", err)
	}
}
