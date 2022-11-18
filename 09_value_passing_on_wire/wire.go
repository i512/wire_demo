//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// Входные параметры инжектятся автоматически
func InitializeWorker(host string) (*Worker, func(), error) {
	wire.Build(NewLogger, DBConnectionProvider, NewWorker)

	return &Worker{}, nil, nil
}

func DBConnectionProvider(host string, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(host, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
