//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// Несколько входных параметров одного типа - недопустимо
func InitializeWorker(logLevel string, host string) (*Worker, func(), error) {
	wire.Build(NewLogger, DBConnectionProvider, NewWorker) // ошибка!

	return &Worker{}, nil, nil
}

func DBConnectionProvider(host string, logger *Logger) (*DBConnection, func(), error) {
	conn, err := NewDBConnection(host, logger)

	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}

// wire: /Users/ilyamorozov/play/wire/08_multiple_value_passing_on_wire/wire.go:10:2: multiple bindings for string
//	current:
//	<- argument host to injector function InitializeWorker (/Users/ilyamorozov/play/wire/08_multiple_value_passing_on_wire/wire.go:9:1)
//	previous:
//	<- argument logLevel to injector function InitializeWorker (/Users/ilyamorozov/play/wire/08_multiple_value_passing_on_wire/wire.go:9:1)
//wire: github.com/i512/wire_demo/08_multiple_value_passing_on_wire: generate failed
//wire: at least one generate failure
