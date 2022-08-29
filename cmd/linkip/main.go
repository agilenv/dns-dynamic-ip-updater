package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Overload()
	if err := buildApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
