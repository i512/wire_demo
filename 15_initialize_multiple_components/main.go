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

	worker2, cleanupDB, err := InitializeWorker2(config)
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()
	defer cleanupDB()

	worker.Work()
	worker2.Work()
}

type Worker2 struct{}

func NewWorker2(logger *Logger, db DBClient) (*Worker2, error) {
	fmt.Println("init worker 2")
	return &Worker2{}, nil
}

func (w *Worker2) Work() {
	fmt.Println("working 2")
}
