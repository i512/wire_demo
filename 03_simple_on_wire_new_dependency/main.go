package main

import "fmt"

func main() {
	worker := InitializeWorker()

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

func NewDBConnection(logger *Logger) *DBConnection {
	fmt.Println("init db")
	return &DBConnection{}
}

func NewLogger() *Logger {
	fmt.Println("init logger")
	return &Logger{}
}

// logger - новая зависимость!
func NewWorker(logger *Logger, db *DBConnection) *Worker {
	fmt.Println("init worker")
	return &Worker{}
}

func (w *Worker) Work() {
	fmt.Println("working")
}
