//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeWorker(config AppConfig) (*Worker, func(), error) {
	wire.Build(
		NewLogger,
		DBConnectionProvider,
		wire.FieldsOf(new(AppConfig), "DBConfig", "LoggerConfig"),
		NewWorker,

		wire.Bind(new(DBClient), new(*DBConnection)),
	)

	return &Worker{}, nil, nil
}

func DBConnectionProvider(config DBConfig, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(config, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
