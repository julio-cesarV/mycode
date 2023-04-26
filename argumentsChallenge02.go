package main

import "os"

func main() {
    // Get the arguments passed in from the command line
    args := os.Args

    // Iterate over the arguments and display them with their position
    for i, arg := range args {
        if i == 0 {
            continue // skip the program name
        }
        println("Argument", i, ":", arg)
    }
}

