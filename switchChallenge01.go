package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {
    switch gove := runtime.Version(); {
    case strings.HasPrefix(gove, "go1.19"), strings.HasPrefix(gove, "go1.18"):
        fmt.Println("You are using the latest version of GoLang")
    case strings.HasPrefix(gove, "go1.17"):
        fmt.Println("This version of Go is fine")
    case strings.HasPrefix(gove, "go1.16"), strings.HasPrefix(gove, "go1.15"):
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:
        fmt.Println("I cannot make a recommendation.")
    }
}

