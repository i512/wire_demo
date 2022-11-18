//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeWorker() (*Worker, error) {
	wire.Build(NewLogger, NewDBConnection, NewWorker)

	return &Worker{}, nil
}
