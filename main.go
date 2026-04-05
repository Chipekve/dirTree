package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    textFileName := "note.txt"
    imageFileName := "utka.jpg"

    f, err := os.Open(textFileName)
    if err != nil {
        log.Fatal(err)
    }

    p, err := os.Open(imageFileName)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    defer p.Close()

    data, err := io.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }

    photo, err := io.ReadAll(p)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("В файле %s находится текст:\n%s\n", textFileName, string(data))
    fmt.Printf("Файл %s занимает %d KB\n", imageFileName, len(photo)/1000)
    fmt.Printf("Файл %s занимает %d byte\n", textFileName, len(data))

}