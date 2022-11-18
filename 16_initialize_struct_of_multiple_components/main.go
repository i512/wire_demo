package main

import (
	"log"
)

func main() {
	config := AppConfig{
		DBConfig{Host: "psql://host:port"},
		LoggerConfig{Level: "info"},
	}

	combo, cleanup, err := Initialize(config)
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	combo.Worker.Work()
	combo.Worker2.Work()
}

type ComboStruct struct {
	*Worker
	*Worker2
}
