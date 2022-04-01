package main

import (
	"chessMoves/models"
	"fmt"
)

const (
	printBoard           = true
	printBlockPiecesInfo = true
	printQueenPosition   = true
	randomQueenPosition  = true
	boardSize            = 8
	randomPieces         = 16
)

func main() {
	// Create an empty board NxN
	board := models.NewBoard(boardSize)

	var queen models.Piece
	if randomQueenPosition {
		// Create own target piece, a Queen (with a random XY position)
		queen = models.Piece{Title: "Queen", Position: *models.NewRandomPosition(boardSize)}
	} else {
		// Create own target piece, a Queen (with a specific XY position)
		queen = models.Piece{Title: "Queen", Position: *models.NewCoordinate(3, 3)}
	}

	if printQueenPosition {
		queen.PrintPosition()
	}

	// Place the queen in our board
	_ = board.SetPiece(queen.Position)

	// Place few other random pieces to block our queen movements
	board.SetRandomPieces(randomPieces)

	// Print board with all pieces placed (0 = empty square, 1 = filled square)
	if printBoard {
		board.Print()
	}

	// Check possible squares that queen can be moved
	possibleMoves := checkPossibleMoves(queen, board)

	// Print obtained results
	printResults(possibleMoves)
}

func printResults(possibleMoves []models.Direction) {
	fmt.Println(len(possibleMoves), "possible moves found")
	for _, movement := range possibleMoves {
		fmt.Println(movement)
	}
}

func checkPossibleMoves(piece models.Piece, board models.Board) []models.Direction {
	var (
		east          = true
		west          = true
		north         = true
		south         = true
		northEast     = true
		northWest     = true
		southEast     = true
		southWest     = true
		possibleMoves []models.Direction
		move          = models.NewDirection(piece)
	)

	// Check all directions at once - O(n)
	for sqr := 1; sqr < len(board[0]); sqr++ {
		possibleMoves = checkDirection(move.ToNorth(sqr), board, &north, possibleMoves)         //look at east, starting from piece
		possibleMoves = checkDirection(move.ToSouth(sqr), board, &south, possibleMoves)         //look at west, starting from piece
		possibleMoves = checkDirection(move.ToEast(sqr), board, &east, possibleMoves)           //look at north, starting from piece
		possibleMoves = checkDirection(move.ToWest(sqr), board, &west, possibleMoves)           //look at south, starting from piece
		possibleMoves = checkDirection(move.ToNorthEast(sqr), board, &northEast, possibleMoves) //look at northEast, starting from piece
		possibleMoves = checkDirection(move.ToNorthWest(sqr), board, &northWest, possibleMoves) //look at northWest, starting from piece
		possibleMoves = checkDirection(move.ToSouthEast(sqr), board, &southEast, possibleMoves) //look at southEast, starting from piece
		possibleMoves = checkDirection(move.ToSouthWest(sqr), board, &southWest, possibleMoves) //look at southWest, starting from piece

		// Validate if it is possible to continue in any direction
		if !(west || east || north || south || northEast || northWest || southEast || southWest) {
			break
		}
	}

	//return possibleMoves
	return sortMoves(possibleMoves) // Return possible movements sorted by main direction
}

func sortMoves(arr []models.Direction) []models.Direction {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if sortByDirections(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func sortByDirections(a models.Direction, b models.Direction) bool {
	return models.Directions[a.Name] > models.Directions[b.Name]
}

func checkDirection(direction models.Direction, board models.Board, canContinue *bool, moves []models.Direction) []models.Direction {
	if *canContinue {
		var (
			x = direction.Coordinate.X
			y = direction.Coordinate.Y
		)

		if coordinatesRemainInBoard(x, y, len(board)) {
			if board[x][y] != 1 {
				return append(moves, models.Direction{Name: direction.Name, Coordinate: models.Coordinate{X: x, Y: y}})
			}
			if printBlockPiecesInfo {
				fmt.Printf("a wild piece was found blocking the queen's movement to the %s at [%d,%d]\n", direction.Name, x, y)
			}
		}
		*canContinue = false
	}
	return moves
}

// coordinatesRemainInBoard returns true if coordinates (x, y) results in a square inside the board
func coordinatesRemainInBoard(x int, y int, tableLen int) bool {
	return x >= 0 && x < tableLen &&
		y >= 0 && y < tableLen
}
