package main

import (
	"fmt"
	"log"
)

func main() {
	worker, cleanup, err := InitializeWorker(
		DBConfig{Host: "psql://host:port"},
		LoggerConfig{Level: "info"},
	)
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	worker.Work()
}

type DBConnection struct{}
type Logger struct{}
type Worker struct{}

type DBConfig struct {
	Host string
}

func NewDBConnection(config DBConfig, logger *Logger) (*DBConnection, error) {
	fmt.Println("connecting to host: ", config.Host)
	return &DBConnection{}, nil
}

func (c *DBConnection) Close() {
	fmt.Println("closing db")
}

type LoggerConfig struct {
	Level string
}

func NewLogger(config LoggerConfig) *Logger {
	fmt.Println("init logger with level: ", config.Level)
	return &Logger{}
}

func NewWorker(logger *Logger, db *DBConnection) (*Worker, error) {
	fmt.Println("init worker")
	return &Worker{}, nil
}

func (w *Worker) Work() {
	fmt.Println("working")
}
