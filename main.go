package main

import (
	"chessMoves/models"
	"fmt"
)

func main() {
	const (
		tableSize    = 8
		randomPieces = 4
	)

	var (
		table models.Table
	)

	table = models.NewTable(tableSize)

	queen := models.Piece{Title: "Queen", Position: *models.NewRandomPosition(tableSize)}
	queen.PrintPosition()

	_ = table.SetPiece(queen.Position)

	// set few random pieces
	table.SetRandomPieces(randomPieces)

	table.PrintTable()

	possibleMoves := checkPossibleMoves(queen, table)

	fmt.Println(len(possibleMoves), "possible moves found")
	for _, move := range possibleMoves {
		fmt.Println(move)
	}
}

func checkPossibleMoves(piece models.Piece, table models.Table) []models.Location {
	var (
		east          = true
		west          = true
		north         = true
		south         = true
		northEast     = true
		northWest     = true
		southEast     = true
		southWest     = true
		possibleMoves []models.Location
	)

	for i := 1; i < len(table[0]); i++ { // Check all moves at once
		possibleMoves = checkDirection(piece.Position.X, piece.Position.Y+i, table, "east", &east, possibleMoves)             //look at east, starting from piece
		possibleMoves = checkDirection(piece.Position.X, piece.Position.Y-i, table, "west", &west, possibleMoves)             //look at west, starting from piece
		possibleMoves = checkDirection(piece.Position.X-i, piece.Position.Y, table, "north", &north, possibleMoves)           //look at north, starting from piece
		possibleMoves = checkDirection(piece.Position.X+i, piece.Position.Y, table, "south", &south, possibleMoves)           //look at south, starting from piece
		possibleMoves = checkDirection(piece.Position.X-i, piece.Position.Y+i, table, "northEast", &northEast, possibleMoves) //look at northEast, starting from piece
		possibleMoves = checkDirection(piece.Position.X-i, piece.Position.Y-i, table, "northWest", &northWest, possibleMoves) //look at northWest, starting from piece
		possibleMoves = checkDirection(piece.Position.X+i, piece.Position.Y+i, table, "southEast", &southEast, possibleMoves) //look at southEast, starting from piece
		possibleMoves = checkDirection(piece.Position.X+i, piece.Position.Y-i, table, "southWest", &southWest, possibleMoves) //look at southWest, starting from piece

		if !(west || east || north || south || northEast || northWest || southEast || southWest) {
			break
		}
	}

	return sortMoves(possibleMoves)
}

func sortMoves(arr []models.Location) []models.Location {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if sortByDirections(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func sortByDirections(a models.Location, b models.Location) bool {
	return models.Directions[a.Direction] > models.Directions[b.Direction]
}

func checkDirection(x, y int, table models.Table, direction string, canContinue *bool, moves []models.Location) []models.Location {
	if x >= 0 && x < len(table[0]) && y >= 0 && y < len(table[0]) && *canContinue {
		if table[x][y] != 1 {
			*canContinue = true
			return append(moves, models.Location{Direction: direction, Position: models.Position{X: x, Y: y}})
		}
		fmt.Printf("a blocking piece was found at [%d,%d]\n", x, y)
	}
	*canContinue = false
	return moves
}
