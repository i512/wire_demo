package main

import (
	"fmt"
	"log"
)

func main() {
	config := AppConfig{
		DBConfig{Host: "psql://host:port"},
		LoggerConfig{Level: "info"},
	}

	worker, cleanup, err := InitializeWorker(config)
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	worker.Work()
}

type Worker struct{}

type DBClient interface {
	Get() string
}

func NewWorker(logger *Logger, db DBClient) (*Worker, error) {
	fmt.Println("init worker")
	return &Worker{}, nil
}

func (w *Worker) Work() {
	fmt.Println("working")
}
