//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func Initialize(config AppConfig) (ComboStruct, func(), error) {
	wire.Build(
		wire.Struct(new(ComboStruct), "*"),
		commonSet,
	)

	return ComboStruct{}, nil, nil
}

var commonSet = wire.NewSet(
	wire.FieldsOf(new(AppConfig), "DBConfig", "LoggerConfig"),
	wire.Bind(new(DBClient), new(*DBConnection)),
	NewLogger,
	DBConnectionProvider,
	NewWorker,
	NewWorker2,
)

// Парсер wire не поймет такое выражение, только определение через переменную
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
