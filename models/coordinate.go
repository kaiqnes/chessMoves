package models

import (
	"math/rand"
	"time"
)

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{
		X: x,
		Y: y,
	}
}

func NewRandomPosition(tableSize int) *Coordinate {
	rand.Seed(time.Now().UnixNano())
	return &Coordinate{
		X: rand.Intn(tableSize),
		Y: rand.Intn(tableSize),
	}
}
