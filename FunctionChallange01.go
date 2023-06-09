package main

import "fmt"

// Author structure
type author struct {
    name   string
    branch string
    books  int
    salary int
}

// Method with a receiver of author type
func (a *author) bookup() {
    a.books++
}

// Method with a receiver of author type
func (a author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

    // Initializing the values
    // of the author structure
    res := author{
        name:   "RZFeeser",
        branch: "Pennsylvania",
        books:  14,
        salary: 34000,
    }

    // Calling the method
    res.bookup()
    res.show()
}

