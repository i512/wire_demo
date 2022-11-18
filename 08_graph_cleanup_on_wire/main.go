package main

import (
	"fmt"
	"log"
)

func main() {
	_, cleanup, err := Initialize()
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()
}

type A struct{}
type B struct{}
type C struct{}
type D struct{}

func NewA() (A, func(), error) {
	fmt.Println("init A")
	return A{}, func() {
			fmt.Println("finish A")
		},
		nil
}

func NewB(a A) (B, func(), error) {
	fmt.Println("init B")
	return B{},
		func() {
			fmt.Println("finish B")
		},
		nil
}

func NewC(a A, b B) (C, func(), error) {
	fmt.Println("init C")
	return C{},
		func() {
			fmt.Println("finish C")
		},
		nil
}

func NewD(a A, b B, c C) (D, func(), error) {
	fmt.Println("init D")
	return D{},
		func() {
			fmt.Println("finish D")
		},
		nil
}
