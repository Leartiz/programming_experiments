package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()

		data := []string{
			"Hello,", "this", "is", "a", "pipe", "example.",
		}

		for _, word := range data {
			_, err := pw.Write([]byte(word + " "))
			if err != nil {
				log.Println("write error:", err)
				return
			}
		}
	}()

	for {
		part := make([]byte, 255)
		partSize, err := pr.Read(part)

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		fmt.Printf("Received (%v): %s\n",
			partSize, string(part))
	}
}
