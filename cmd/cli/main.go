package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if err := buildApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
