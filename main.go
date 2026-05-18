package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func printDirRecursive(path, indent string, out io.Writer) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	entriesSort(entries)

	for i, e := range entries {
		connector := "├───"
		if i == len(entries)-1 {
			connector = "└───"
		}
		if e.Name() == ".git" {
			continue
		}

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
		fmt.Fprintf(out, "%s%s %-10s %s %-5s\n", indent, connector, e.Name(), t, size)
		if e.IsDir() {
			printDirRecursive(path+"/"+e.Name(), indent+"│\t", out)
		}
	}
}

func readDir(path, indent string, out io.Writer) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	entriesSort(entries)

	for i, e := range entries {
		fmt.Println(i)
		if e.Name() == ".git" {
			continue
		}

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
		fmt.Fprintf(out, "%s %-10s %-5s %s\n", indent, e.Name(), t, size)
	}
}

func entriesSort(entries []os.DirEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	path := flag.String("path", ".", "default")
	r := flag.Bool("r", false, "recursion")
	flag.Parse()
	if *r {
		printDirRecursive(*path, "", f)
		printDirRecursive(*path, "", os.Stdout)
	} else {
		readDir(*path, "", f)
		readDir(*path, "", os.Stdout)
	}
}
