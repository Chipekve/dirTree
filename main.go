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

        size := info.Size() / 1024
        var t string

        if e.IsDir() {
            t = "dir"
        } else {
            t = "file"
        }

        fmt.Printf("%s %s %d KB\n", e.Name(), t, size)
    }
}