package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AbhishekSinghDev/scaleURL/internal/config"
	"github.com/AbhishekSinghDev/scaleURL/internal/features/url"
	"github.com/AbhishekSinghDev/scaleURL/internal/logger"
	"github.com/AbhishekSinghDev/scaleURL/internal/middleware"
	"github.com/AbhishekSinghDev/scaleURL/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	config.MustLoad()
	logger.Init()
	cfg := config.Get()

	pool := postgres.New(cfg)
	router := gin.New()

	urlRepo := url.NewRepository(pool.Db)
	urlService := url.NewService(urlRepo)
	urlHandler := url.NewHandler(urlService)

	router.POST("/shorten", urlHandler.Create)

	// middlewares
	router.Use(middleware.RequestIdMiddleware(), middleware.LoggerMiddleware(), gin.Recovery())

	address := fmt.Sprintf("localhost:%s", cfg.Port)
	server := &http.Server{
		Addr: address,
		Handler: router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Info().Str("addr", address).Msg("starting server")
	// start the server in goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
    		log.Error().Err(err).Msg("failed to start server")
    		os.Exit(1)
		}
	}()

	// wait for syscall signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer pool.Db.Close()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Str("error", err.Error()).Msg("force shutdown")
	}
}
