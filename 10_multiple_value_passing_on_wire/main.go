package main

import (
	"fmt"
	"log"
)

func main() {
	worker, cleanup, err := InitializeWorker("info", "psql://host:port")
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

func (c *DBConnection) Close() {
	fmt.Println("closing db")
}

// Хотим указать уровень логирования - новый параметр тоже строка
func NewLogger(logLevel string) *Logger {
	fmt.Println("init logger with level: ", logLevel)
	return &Logger{}
}

func NewWorker(logger *Logger, db *DBConnection) (*Worker, error) {
	fmt.Println("init worker")
	return &Worker{}, nil
}

func (w *Worker) Work() {
	fmt.Println("working")
}
