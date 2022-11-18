package main

import "fmt"

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
