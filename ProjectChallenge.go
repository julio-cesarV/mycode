package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"

)

func main() {
    var answer int
    var correctCount int
    var incorrectCount int


    file, err := os.Open("Questions.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()



    reader := csv.NewReader(file)
    data, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }



    fmt.Println("Welcome to the Math Quiz for ALTA3 Project/Challenge :) !!!")

    for _, row := range data {
        question := row[0]
        answerStr := row[1]
        correctAnswer, err := strconv.Atoi(answerStr)
        if err != nil {
            fmt.Println("Error :", err)
            return
        }

        fmt.Println(question)
        fmt.Scanln(&answer)
  
        if answer == correctAnswer {
            fmt.Println("Correct!  :) ")
            correctCount++
        } else {
            fmt.Println("Incorrect. :( ")
        }
    }

    incorrectCount = len(data) - correctCount

    fmt.Printf("\nYou got %d out of 10 questions correct.\n", correctCount)
    fmt.Printf("\nYou got %d incorrect answers\n", incorrectCount)
    fmt.Println("\nThanks for answer this quiz, GOOD BYE !!!\n")
}

