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

type Direction struct {
	Name       string
	Coordinate Coordinate
}

func (c *Coordinate) SetRandomPosition(tableSize int) {
	rand.Seed(time.Now().UnixNano())
	c.X = rand.Intn(tableSize)
	c.Y = rand.Intn(tableSize)
}

func NewDirection(piece Piece) *Direction {
	return &Direction{Coordinate: *NewCoordinate(piece.Position.X, piece.Position.Y)}
}

func (d *Direction) ToNorth(squares int) Direction {
	return Direction{
		Name: "north",
		Coordinate: Coordinate{
			X: d.Coordinate.X - squares,
			Y: d.Coordinate.Y,
		},
	}
}

func (d *Direction) ToSouth(squares int) Direction {
	return Direction{
		Name: "south",
		Coordinate: Coordinate{
			X: d.Coordinate.X + squares,
			Y: d.Coordinate.Y,
		},
	}
}

func (d *Direction) ToEast(squares int) Direction {
	return Direction{
		Name: "east",
		Coordinate: Coordinate{
			X: d.Coordinate.X,
			Y: d.Coordinate.Y + squares,
		},
	}
}

func (d *Direction) ToWest(squares int) Direction {
	return Direction{
		Name: "west",
		Coordinate: Coordinate{
			X: d.Coordinate.X,
			Y: d.Coordinate.Y - squares,
		},
	}
}

func (d *Direction) ToNorthWest(squares int) Direction {
	return Direction{
		Name: "northWest",
		Coordinate: Coordinate{
			X: d.Coordinate.X - squares,
			Y: d.Coordinate.Y - squares,
		},
	}
}

func (d *Direction) ToSouthWest(squares int) Direction {
	return Direction{
		Name: "southWest",
		Coordinate: Coordinate{
			X: d.Coordinate.X + squares,
			Y: d.Coordinate.Y - squares,
		},
	}
}

func (d *Direction) ToNorthEast(squares int) Direction {
	return Direction{
		Name: "northEast",
		Coordinate: Coordinate{
			X: d.Coordinate.X - squares,
			Y: d.Coordinate.Y + squares,
		},
	}
}

func (d *Direction) ToSouthEast(squares int) Direction {
	return Direction{
		Name: "southEast",
		Coordinate: Coordinate{
			X: d.Coordinate.X + squares,
			Y: d.Coordinate.Y + squares,
		},
	}
}
