package main

import (
	"fmt"
	"log"
)

func main() {
	worker, cleanup, err := InitializeWorker()
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

func NewDBConnection(logger *Logger) (*DBConnection, error) {
	fmt.Println("init db")
	return &DBConnection{}, nil
}

// Нужно закрыть коннект к базе
func (c *DBConnection) Close() {
	fmt.Println("closing db")
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
