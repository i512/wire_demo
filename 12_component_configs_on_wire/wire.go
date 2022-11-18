//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeWorker(_ DBConfig, _ LoggerConfig) (*Worker, func(), error) {
	wire.Build(NewLogger, DBConnectionProvider, NewWorker)

	return &Worker{}, nil, nil
}

// Кастомный провайдер для логгера больше не нужен

func DBConnectionProvider(config DBConfig, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(config, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
