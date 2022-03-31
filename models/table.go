package models

import (
	"errors"
	"fmt"
)

type Table [][]int

func NewTable(size int) (table Table) {
	for i := 0; i < size; i++ {
		table = append(table, make([]int, size))
	}
	return table
}

func (t Table) SetPiece(pos Position) error {
	if t[pos.X][pos.Y] == 0 {
		t[pos.X][pos.Y] = 1
		return nil
	}
	return errors.New("position already is occupied")
}

func (t Table) PrintTable() {
	for i := 0; i < len(t); i++ {
		fmt.Printf("[ ")
		for j := 0; j < len(t); j++ {
			fmt.Printf("%d", t[i][j])
			if j < len(t)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("] %d\n", i)
	}
	for i := 0; i < len(t); i++ {
		fmt.Printf("  %d", i)
	}
	fmt.Println()
}

func (t Table) SetRandomPieces(qtd int) {
	for i := 0; i < qtd; i++ {
		err := t.SetPiece(*NewRandomPosition(len(t)))
		if err != nil {
			i--
		}
	}
}
