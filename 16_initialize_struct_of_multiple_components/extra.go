package main

import "fmt"

type Worker2 struct{}

func NewWorker2(logger *Logger, db DBClient) (*Worker2, error) {
	fmt.Println("init worker 2")
	return &Worker2{}, nil
}

func (w *Worker2) Work() {
	fmt.Println("working 2")
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

type DBConnection struct{}
type Logger struct{}

type AppConfig struct {
	DBConfig
	LoggerConfig
}

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

func (c *DBConnection) Get() string {
	return "string"
}

type LoggerConfig struct {
	Level string
}

func NewLogger(config LoggerConfig) *Logger {
	fmt.Println("init logger with level: ", config.Level)
	return &Logger{}
}
