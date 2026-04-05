package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}

		var t string
		var size string

		if e.IsDir() {
			size = "-"
		} else {
			size = fmt.Sprintf("%d KB", info.Size()/1024)
		}
		if e.IsDir() {
			t = "dir"
		} else {
			t = "file"
		}
		fmt.Printf("%-10s %-5s %s\n", e.Name(), t, size)
	}
}
