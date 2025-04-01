package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type config struct {
	Env               string
	Location          string
	Port              string
	MaxConnection     int
	ReadTimeout       int
	WriteTimeout      int
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
	MaxIdleConnection int
	MaxOpenConnection int
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

	c.DBHost = viper.GetString("DB_HOST")
	if c.DBHost == "" {
		slog.Error("DB_HOST is required")
		result = false
	}

	dbPort := viper.GetString("DB_PORT")
	if dbPort == "" {
		slog.Error("DB_PORT is required")
		result = false
	}

	dp, err := strconv.Atoi(dbPort)
	if err != nil {
		slog.Error("invalid DB_PORT value")
		result = false
	}
	c.DBPort = dp

	c.DBUser = viper.GetString("DB_USER")
	if c.DBUser == "" {
		slog.Error("DB_USER is required")
		result = false
	}

	c.DBPassword = viper.GetString("DB_PASSWORD")
	if c.DBPassword == "" {
		slog.Error("DB_PASSWORD is required")
		result = false
	}

	c.DBName = viper.GetString("DB_NAME")
	if c.DBName == "" {
		slog.Error("DB_NAME is required")
		result = false
	}

	c.MaxIdleConnection = viper.GetInt("MAX_IDLE_CONNECTION")
	if c.MaxIdleConnection == 0 {
		slog.Error("MAX_IDLE_CONNECTION is required")
		result = false
	}

	c.MaxOpenConnection = viper.GetInt("MAX_OPEN_CONNECTION")
	if c.MaxOpenConnection == 0 {
		slog.Error("MAX_OPEN_CONNECTION is required")
		result = false
	}

	return result
}
