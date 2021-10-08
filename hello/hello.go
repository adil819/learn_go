package main

import (
    "fmt"
    "log"

    "GO/greetings"
)

func main() {
    // Set properties of the predifined logger, including 
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Miftah", "Bagus", "Yuan"}

    for idk, name := range names {
        fmt.Println("Read : ", name, idk)
    }

    // get a greeting message and print it
    message, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}
