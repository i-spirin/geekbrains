package main

import (
	"log"

	"github.com/i-spirin/geekbrains/lesson_08/config"
)

func main() {
	config := config.Config{}
	config.Parse()

	log.Printf("Parsed config: %+v", config)
}
