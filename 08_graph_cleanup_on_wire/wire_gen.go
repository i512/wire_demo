// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func Initialize() (D, func(), error) {
	a, cleanup, err := NewA()
	if err != nil {
		return D{}, nil, err
	}
	b, cleanup2, err := NewB(a)
	if err != nil {
		cleanup()
		return D{}, nil, err
	}
	c, cleanup3, err := NewC(a, b)
	if err != nil {
		cleanup2()
		cleanup()
		return D{}, nil, err
	}
	d, cleanup4, err := NewD(a, b, c)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return D{}, nil, err
	}
	return d, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
