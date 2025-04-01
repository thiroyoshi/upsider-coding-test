package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"upsider-coding-test/cmd/api/controller/invoices"
	ul "upsider-coding-test/internal/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/netutil"
	"golang.org/x/sync/errgroup"
)

const (
	PREFIX = "/api"
)

var ErrFailedToStartServer = errors.New("failed to start server")

func getRouteHandler() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	// Get handlers
	gc := invoices.NewGetController()
	pc := invoices.NewPostController()

	// Set routes
	v1 := e.Group(PREFIX)
	{
		v1.POST("/invoices", pc.Post())
		v1.GET("/invoices", gc.Get())
	}

	return e
}

func main() {
	slog.Info("start initializing server.")

	// initialize logger
	logger := slog.New(ul.NewHandler())
	slog.SetDefault(logger)

	// configurations
	cfg := config{}
	if result := cfg.readAndValidate(); !result {
		slog.Error("failed to start server: invalid config")
		return
	}

	// set conditions
	svMng := http.Server{
		Handler:      getRouteHandler(),
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}

	// set Port
	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		errMessage := fmt.Sprintf("failed to listen on port %s: %v", cfg.Port, err)
		slog.Error(errMessage)
		return
	}

	// initialize listener
	listener := netutil.LimitListener(ln, cfg.MaxConnection)
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			slog.Error("failed to close listener: ", "error", err)
		}
	}(listener)

	slog.Info("set server configs.")

	// start server
	var eg errgroup.Group
	eg.Go(func() error {
		err := svMng.Serve(listener)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrFailedToStartServer, err)
		}
		return nil
	})

	slog.Info("started server on :" + cfg.Port)

	if err := eg.Wait(); err != nil {
		slog.Error("error running server: ", "error", err)
	}
}
