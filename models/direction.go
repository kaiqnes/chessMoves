package models

import (
	"math/rand"
	"time"
)

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
	NORTHEAST
	SOUTHEAST
	NORTHWEST
	SOUTHWEST
)

var (
	Directions = map[string]int{
		"east":      EAST,
		"west":      WEST,
		"north":     NORTH,
		"south":     SOUTH,
		"northEast": NORTHEAST,
		"northWest": NORTHWEST,
		"southEast": SOUTHEAST,
		"southWest": SOUTHWEST,
	}
)

type Position struct {
	X, Y int
}

func NewRandomPosition(tableSize int) *Position {
	rand.Seed(time.Now().UnixNano())
	return &Position{
		X: rand.Intn(tableSize),
		Y: rand.Intn(tableSize),
	}
}

func (p *Position) SetRandomPosition(tableSize int) {
	rand.Seed(time.Now().UnixNano())
	p.X = rand.Intn(tableSize)
	p.Y = rand.Intn(tableSize)
}

type Location struct {
	Direction string
	Position  Position
}
