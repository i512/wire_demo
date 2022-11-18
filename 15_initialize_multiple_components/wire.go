//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeWorker(config AppConfig) (*Worker, func(), error) {
	wire.Build(commonSet)

	return &Worker{}, nil, nil
}

func InitializeWorker2(config AppConfig) (*Worker2, func(), error) {
	wire.Build(commonSet)

	return &Worker2{}, nil, nil
}

var commonSet = wire.NewSet(
	wire.FieldsOf(new(AppConfig), "DBConfig", "LoggerConfig"),
	wire.Bind(new(DBClient), new(*DBConnection)),
	NewLogger,
	DBConnectionProvider,
	NewWorker,
	NewWorker2,
)

// Парсер wire не поймет такое выражение
func _commonSet() wire.ProviderSet {
	return wire.NewSet(
	// ...
	)
}

func DBConnectionProvider(config DBConfig, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(config, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
