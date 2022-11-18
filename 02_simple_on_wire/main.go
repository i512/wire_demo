package main

import "fmt"

func main() {
	worker := InitializeWorker()

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

// Dependency providers:

func NewDBConnection(logger *Logger) *DBConnection {
	fmt.Println("init db")
	return &DBConnection{}
}

func NewLogger() *Logger {
	fmt.Println("init logger")
	return &Logger{}
}

func NewWorker(db *DBConnection) *Worker {
	fmt.Println("init worker")
	return &Worker{}
}

// =====================

func (w *Worker) Work() {
	fmt.Println("working")
}
