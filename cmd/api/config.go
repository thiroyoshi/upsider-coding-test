package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type config struct {
	ProjectID                 string
	ProjectNumber             string
	Env                       string
	Location                  string
	Port                      string
	MaxConnection             int
	ReadTimeout               int
	WriteTimeout              int
	MemorystoreAddress        string
	MemorystorePort           string
	MemorystoreMaxIdle        int
	MemorystoreMaxActive      int
	MemorystoreIdleTimeoutSec int
	SpannerInstanceID         string
	SpannerDatabaseID         string
	SpannerMinOpened          uint64
	SpannerMaxOpened          uint64
	SpannerMaxIdle            uint64
	EndpointHash              string
	WebEndpoint               string
	PubsubTopicMail           string
	GCPAPIKey                 string
}

func (c *config) readAndValidate() bool {
	// 環境変数を自動的に読み取る
	viper.AutomaticEnv()

	// 設定ファイルのパスを追加
	viper.AddConfigPath("./viper/")

	viper.SetConfigType("yaml")

	viper.SetConfigName(os.Getenv("ENV"))

	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("failed to read config file", "error", err)
		return false
	}

	result := true

	c.Env = viper.GetString("ENV")
	if c.Env == "" {
		slog.Error("ENV is required")
		result = false
	}

	c.Port = viper.GetString("LISTEN_PORT")
	if c.Port == "" {
		slog.Error("LISTEN_PORT is required")
		result = false
	}
	port, err := strconv.Atoi(c.Port)
	if err != nil {
		slog.Error("invalid LISTEN_PORT value")
		result = false
	}

	if port < 1 || port > 65535 {
		slog.Error("invalid LISTEN_PORT value, port value out of range: ", "port", port)
		result = false
	}

	c.MaxConnection = viper.GetInt("MAX_CONNECTION")
	if c.MaxConnection == 0 {
		slog.Error("MAX_CONNECTION is required")
		result = false
	}

	c.ReadTimeout = viper.GetInt("CONNECTION_TIMEOUT_READ")
	if c.ReadTimeout == 0 {
		slog.Error("CONNECTION_TIMEOUT_READ is required")
		result = false
	}

	c.WriteTimeout = viper.GetInt("CONNECTION_TIMEOUT_WRITE")
	if c.WriteTimeout == 0 {
		slog.Error("CONNECTION_TIMEOUT_WRITE is required")
		result = false
	}

	return result
}
