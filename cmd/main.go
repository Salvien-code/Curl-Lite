// Sample program to show how to write a simple version of curl using Go.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ./main <url>")
	}
}

func main() {
	// Get a response from the web server.
	res, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copies from the Body to Stdout.
	io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		log.Fatalln(err)
	}
}
