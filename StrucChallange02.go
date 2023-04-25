package main

import "fmt"

type Astro struct {
    Name     string
    Age      int
    Mission  string
    IsNeeded bool
}
func main() {
    astro1 := Astro{Name: "John Smith", Age: 35, Mission: "Apollo 11", IsNeeded: true}
    astro2 := Astro{Name: "Jane Doe", Age: 29, Mission: "Mars One", IsNeeded: false}
    astro3 := Astro{Name: "Bob Johnson", Age: 42, Mission: "Jupiter Probe", IsNeeded: true}

    fmt.Println("Astro 1:", astro1)
    fmt.Println("Astro 2:", astro2)
    fmt.Println("Astro 3:", astro3)

    p := []Astro{astro1, astro2, astro3}

    fmt.Println(p)

    fmt.Println(p[2].Mission)
}
