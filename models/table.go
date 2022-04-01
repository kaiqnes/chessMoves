package models

import (
	"errors"
	"fmt"
)

type Board [][]int

func NewBoard(size int) (table Board) {
	for i := 0; i < size; i++ {
		table = append(table, make([]int, size))
	}
	return table
}

func (b Board) SetPiece(pos Coordinate) error {
	if b[pos.X][pos.Y] == 0 {
		b[pos.X][pos.Y] = 1
		return nil
	}
	return errors.New("square already is occupied")
}

func (b Board) Print() {
	for i := 0; i < len(b); i++ {
		fmt.Printf("[ ")
		for j := 0; j < len(b); j++ {
			fmt.Printf("%d", b[i][j])
			if j < len(b)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("] %d\n", i)
	}
	for i := 0; i < len(b); i++ {
		fmt.Printf("  %d", i)
	}
	fmt.Println()
}

func (b Board) SetRandomPieces(qtd int) {
	for i := 0; i < qtd; i++ {
		err := b.SetPiece(*NewRandomPosition(len(b)))
		if err != nil {
			i--
		}
	}
}
