package main

import "fmt"

type Direction int

const (
	Left  Direction = -1
	Right Direction = 1
)

type Move struct {
	Direction Direction
	Letter    rune // use rune(-1) to represent nothing captured
}

type Moves []Move

type Board []rune

func (move Move) String() string {
	directionName := "LEFT"
	if move.Direction == Right {
		directionName = "Right"
	}

	letterName := string(move.Letter)
	if move.Letter == -1 {
		letterName = "nil"
	}

	return fmt.Sprintf("%s:%s", directionName, letterName)
}
