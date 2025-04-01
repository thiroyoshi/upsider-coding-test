package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"upsider-coding-test/cmd/api/controller/invoices"
	auth "upsider-coding-test/internal/auth"
	ul "upsider-coding-test/internal/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/netutil"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

const (
	PREFIX = "/api"
)

var ErrFailedToStartServer = errors.New("failed to start server")

func getRouteHandler(db *gorm.DB) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	// set middleware of auth api key
	e.Use(auth.AuthMiddleware(db))

	// Get handlers
	gc := invoices.NewGetController(db)
	pc := invoices.NewPostController(db)

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

	// initialize database
	db, err := initDB(cfg)
	if err != nil {
		slog.Error("failed to start server: failed to initialize database")
		return
	}

	// set conditions
	svMng := http.Server{
		Handler:      getRouteHandler(db),
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

func initDB(cfg config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	gormConfig := &gorm.Config{
		Logger: gl.Default.LogMode(gl.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to get database instance", "error", err)
		return nil, err
	}

	// コネクションプールの設定
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnection)

	return db, nil
}
