package models

import "fmt"

type Piece struct {
	Title    string
	Position Coordinate
}

func (p *Piece) PrintPosition() {
	fmt.Printf("%s is placed in row %d column %d\n", p.Title, p.Position.X, p.Position.Y)
}
