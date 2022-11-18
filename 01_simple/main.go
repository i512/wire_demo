package main

import "fmt"

func main() {
	logger := NewLogger()
	db := NewDBConnection(logger)
	worker := NewWorker(db)

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

func NewDBConnection(logger *Logger) *DBConnection {
	return &DBConnection{}
}

func NewLogger() *Logger {
	return &Logger{}
}

func NewWorker(db *DBConnection) *Worker {
	return &Worker{}
}

func (w *Worker) Work() {
	fmt.Println("working")
}
