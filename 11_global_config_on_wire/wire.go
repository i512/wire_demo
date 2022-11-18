//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeWorker(config Config) (*Worker, func(), error) {
	wire.Build(LoggerProvider, DBConnectionProvider, NewWorker)

	return &Worker{}, nil, nil
}

// Нам придется написать кастомные провайдеры чтобы достать параметры из конфига
func LoggerProvider(config Config) *Logger {
	logger := NewLogger(config.LogLevel)

	return logger
}

func DBConnectionProvider(config Config, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(config.DBHost, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
