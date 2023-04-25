package main

import "fmt"

func main() {
    // map extensions
    languageExtensions := map[string]string{
        "Python":   ".py",
        "Golang":   ".go",
        "Java":     ".java",
        "Ansible":  ".yml",
        "C++":      ".cpp",
    }

    // display the map
    fmt.Println("File extensions:", languageExtensions)

    // remove C++
    delete(languageExtensions, "C++")

    // add Julia
    languageExtensions["Julia"] = ".jl"

    // display the updated map
    fmt.Println("File extensions updated:", languageExtensions)
}

