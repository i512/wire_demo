//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeWorker() (*Worker, func(), error) {
	wire.Build(NewLogger, DBConnectionProvider, NewWorker)

	return &Worker{}, nil, nil
}

func DBConnectionProvider(logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(logger)
	// Каждый провайдер может вернуть cleanup функцию
	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
