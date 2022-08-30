package main

import (
	"log"
	"os"
)

func main() {
	if err := buildApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
