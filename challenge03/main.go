package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing argument: filename")
	}

	filename := os.Args[1]

	if filename == "" {
		log.Fatal("Filename must be specified.")
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
	fmt.Println()
}
