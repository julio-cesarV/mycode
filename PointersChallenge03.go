package main

import "fmt"

type Player struct {
    Lives     int
    Stage     int
    Inventory []string
}

func (p *Player) GreenMushroom() {
    p.Lives++
}

func (p *Player) Pickup(item string) {
    p.Inventory = append(p.Inventory, item)
}

func (p *Player) CanWhistle() bool {
    return p.Stage >= 5
}

func main() {
    mario := Player{
        Lives:     3,
        Stage:     1,
        Inventory: []string{"mushroom"},
    }

    mario.GreenMushroom()
    mario.Pickup("flower")
    mario.Stage = 4
    canWhistle := mario.CanWhistle()
    fmt.Printf("Mario: %+v, CanWhistle: %v\n", mario, canWhistle)
}

