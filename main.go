// main package
package main

import (
    "fmt"
    "mygame/models"
)

func main() {
    mario := models.Player{
        Lives:     3,
        Stage:     1,
        Inventory: []string{"mushroom"},
    }

    // Mario touches a green mushroom and gains a life
    mario.GreenMushroom()

    // Mario picks up a fire flower
    mario.Pickup("fire flower")

    // Mario checks if he can whistle
    canWhistle := mario.CanWhistle()

    fmt.Printf("Mario: %+v, CanWhistle: %v\n", mario, canWhistle)
}

