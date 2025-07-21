package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vandannandwana/Basic-E-Commerce/internal/config"
	"github.com/vandannandwana/Basic-E-Commerce/internal/http/handlers/product"
	"github.com/vandannandwana/Basic-E-Commerce/internal/storage/sqlite"
)

func main() {

	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version: ", "1.0.0"))

	router := http.NewServeMux()

	router.HandleFunc("POST /api/products", product.New(storage))
	router.HandleFunc("GET /api/products/{id}", product.GetProductById(storage))
	router.HandleFunc("GET /api/products", product.GetProducts(storage))
	router.HandleFunc("DELETE /api/products/{id}", product.DeleteProductById(storage))

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server Started ", slog.String("on: ", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error: ", err.Error()))
	}

	slog.Info("Server Shutdown Successfully")

}
