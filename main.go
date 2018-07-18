package main

import (
	"log"
	"os"

	"github.com/st3fan/xliff"
)

func main() {
	doc, err := xliff.FromFile(os.Args[1])
	if err != nil {
		log.Fatal("Failed to read file: ", err)
	}

	if errors := doc.Validate(); errors != nil {
		for _, err := range errors {
			log.Println("Validation error:", err)
		}
		os.Exit(1)
	}
}
