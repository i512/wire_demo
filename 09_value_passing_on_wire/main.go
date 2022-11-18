package main

import (
	"fmt"
	"log"
)

func main() {
	worker, cleanup, err := InitializeWorker("psql://host:port")
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

func NewDBConnection(host string, logger *Logger) (*DBConnection, error) {
	fmt.Println("connecting to host: ", host)
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
