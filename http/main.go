package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type LogWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatal(err)
	}

	lw := LogWriter{}

	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
