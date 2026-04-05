package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func printDirRecursive(path, indent string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}

		t := "file"
		size := fmt.Sprintf("(%db)", info.Size()/1024)

		if e.IsDir() {
			t = "dir"
			size = "-"
		}
		fmt.Printf("%s %-10s %s %-5s\n", indent, e.Name(), t, size)
		if e.IsDir() {
			printDirRecursive(path+"/"+e.Name(), indent+"  /")
		}
	}
}

func readDir(path, indent string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}
		t := "file"
		size := fmt.Sprintf("(%db)", info.Size())

		if e.IsDir() {
			t = "dir"
			size = "-"
		}
		fmt.Printf("%s %-10s %-5s %s\n", indent, e.Name(), t, size)
	}
}

func main() {
	path := flag.String("path", ".", "default")
	r := flag.Bool("r", false, "recursion")
	flag.Parse()
	if *r {
		printDirRecursive(*path, "")
	} else {
		readDir(*path, "")
	}
}
