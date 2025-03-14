package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jhrick/go-template-api/internal/core/infra/config"
	"github.com/jhrick/go-template-api/internal/core/infra/db/postgres"
)

func main() {
	config := config.New()

	ctx := context.Background()

	database := postgres.New(ctx, *config.DB)
	if err := database.Connect(); err != nil {
		panic(err)
	}

	go func() {
		if err := http.ListenAndServe(config.Server.Port, config.Server.Handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	log.Println("Server Running!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Server Closed.")
}
