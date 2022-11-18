//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func Initialize() (D, func(), error) {
	wire.Build(NewA, NewB, NewC, NewD)

	return D{}, nil, nil
}
