package main

import (
	"fmt"
	"log"
)

func main() {
	worker, err := InitializeWorker()
	if err != nil {
		log.Fatal(err)
	}

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

func NewDBConnection(logger *Logger) (*DBConnection, error) {
	fmt.Println("init db")
	return &DBConnection{}, nil
}

func NewLogger() *Logger {
	fmt.Println("init logger")
	return &Logger{}
}

func NewWorker(logger *Logger, db *DBConnection) (*Worker, error) {
	fmt.Println("init worker")
	return &Worker{}, nil
}

func (w *Worker) Work() {
	fmt.Println("working")
}
