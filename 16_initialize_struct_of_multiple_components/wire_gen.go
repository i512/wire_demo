// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func Initialize(config AppConfig) (ComboStruct, func(), error) {
	loggerConfig := config.LoggerConfig
	logger := NewLogger(loggerConfig)
	dbConfig := config.DBConfig
	dbConnection, cleanup, err := DBConnectionProvider(dbConfig, logger)
	if err != nil {
		return ComboStruct{}, nil, err
	}
	worker, err := NewWorker(logger, dbConnection)
	if err != nil {
		cleanup()
		return ComboStruct{}, nil, err
	}
	worker2, err := NewWorker2(logger, dbConnection)
	if err != nil {
		cleanup()
		return ComboStruct{}, nil, err
	}
	comboStruct := ComboStruct{
		Worker:  worker,
		Worker2: worker2,
	}
	return comboStruct, func() {
		cleanup()
	}, nil
}

// wire.go:

var commonSet = wire.NewSet(wire.FieldsOf(new(AppConfig), "DBConfig", "LoggerConfig"), wire.Bind(new(DBClient), new(*DBConnection)), NewLogger,
	DBConnectionProvider,
	NewWorker,
	NewWorker2,
)

// Парсер wire не поймет такое выражение, только определение через переменную
func _commonSet() wire.ProviderSet {
	return wire.NewSet()
}

func DBConnectionProvider(config DBConfig, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(config, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}
