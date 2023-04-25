package main

import "fmt"

type Astro struct {
    Name     string
    Age      int
    Mission  string
    IsNeeded bool
}

type nasaMission struct {
    People  []Astro
    Number  int
    Message string
}

func main() {
    p := []Astro{
        {Name: "John Smith", Age: 35, Mission: "Apollo 11", IsNeeded: true},
        {Name: "Jane Doe", Age: 29, Mission: "Mars One", IsNeeded: false},
        {Name: "Bob Johnson", Age: 42, Mission: "Jupiter Probe", IsNeeded: true},
    }

    s := nasaMission{People: p, Number: 3, Message: "Success"}

    fmt.Println("Full nasaMission struct:", s)

    fmt.Println("People on mission:", s.People)
    fmt.Println("Number of people on mission:", s.Number)
    fmt.Println("Mission message:", s.Message)
}

