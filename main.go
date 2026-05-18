package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func printDirRecursive(path, indent string, out io.Writer, printFiles bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}
	entries = filter(entries)

	if !printFiles {
		var dirsOnly []os.DirEntry
		for _, e := range entries {
			if e.IsDir() {
				dirsOnly = append(dirsOnly, e)
			}
		}
		entries = dirsOnly
	}

	entriesSort(entries)

	for i, e := range entries {
		connector := "├───"
		if i == len(entries)-1 {
			connector = "└───"
		}

		size := ""
		if !e.IsDir() {
			info, err := e.Info()
			if err != nil {
				continue
			}
			if info.Size() == 0 {
				size = " (empty)"
			} else {
				size = fmt.Sprintf(" (%db)", info.Size())
			}
		}

		fmt.Fprintf(out, "%s%s%s%s\n", indent, connector, e.Name(), size)

		if e.IsDir() {
			if i == len(entries)-1 {
				printDirRecursive(path+"/"+e.Name(), indent+"    ", out, printFiles)
			} else {
				printDirRecursive(path+"/"+e.Name(), indent+"│   ", out, printFiles)
			}
		}
	}
}

func filter(entries []os.DirEntry) []os.DirEntry {
	var result []os.DirEntry
	for _, e := range entries {
		if e.Name()[0] != '.' {
			result = append(result, e)
		}
	}
	return result
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

	printDirRecursive(*path, "", f, *r)
	printDirRecursive(*path, "", os.Stdout, *r)
}
