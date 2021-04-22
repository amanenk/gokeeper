package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fdistorted/gokeeper/config"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/validator"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config %+v\n", err)
	}

	err = logger.Load()
	if err != nil {
		log.Fatalf("failed to load logger %+v\n", err)
	}

	err = validator.Load()
	if err != nil {
		logger.Get().Fatal("failed to load logger", zap.Error(err))
	}

	err = database.Load(cfg)
	if err != nil {
		logger.Get().Fatal("failed to load database", zap.Error(err))
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: handlers.NewRouter(),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Get().Info("Listening...", zap.String("listen_url", addr))
		err = server.ListenAndServe()
		if err != nil {
			logger.Get().Error("Failed to initialize HTTP server", zap.Error(err))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
