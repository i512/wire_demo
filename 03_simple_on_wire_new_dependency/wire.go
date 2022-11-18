//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeWorker() *Worker {
	wire.Build(NewLogger, NewDBConnection, NewWorker)

	return &Worker{}
}
