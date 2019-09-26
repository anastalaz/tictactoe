package tictactoe

import "fmt"

// Board represents a single Tic Tac Toe board.
type Board [3][3]int

// Move represents a new move by a player. The first
// two integers represent the coordinates and the third
// the player.
type Move [3]int

// NewBoard returns a new Tic Tac Toe board.
func NewBoard() Board {
	var b Board
	return b
}

// Render prints the Tic Tac Toe board to the console.
func (b Board) Render() {
	fmt.Printf("%d | %d | %d \n", b[0][0], b[0][1], b[0][2])
	fmt.Printf("%d | %d | %d \n", b[1][0], b[1][1], b[1][2])
	fmt.Printf("%d | %d | %d \n", b[2][0], b[2][1], b[2][2])
}

// NewMove returns a Tic Tac Toe board with the new move.
func (b Board) NewMove(m Move) Board {
	b[m[0]][m[1]] = m[2]
	return b
}

// GetWinner returns 1 if the winner is player 1, 2 if the winner
// is player 2 or 0 if the game is a draw.
func (b Board) GetWinner() int {

	for i := 0; i < 3; i++ {
		// check rows
		if b[0][i] == b[1][i] && b[0][i] == b[2][i] && b[0][i] != 0 {
			return b[0][i]
		}
		// check columns
		if b[i][0] == b[i][1] && b[i][0] == b[i][2] && b[i][0] != 0 {
			return b[i][0]
		}
	}

	// Check diagonals
	if b[0][0] == b[1][1] && b[0][0] == b[2][2] && b[0][0] != 0 {
		return b[0][0]
	} else if b[2][0] == b[1][1] && b[2][0] == b[0][2] && b[2][0] != 0 {
		return b[2][0]
	}

	return 0
}

// IsFull returns true if the board is full else false.
func (b Board) IsFull() bool {
	for _, array := range b {
		for _, cell := range array {
			if cell == 0 {
				return false
			}
		}

	}
	return true
}

// LegalMoves returns a slice of Move with the coordinates
// of all the cells that are free. The third parameter of
// Move is always zero.
func (b Board) LegalMoves() (m []Move) {
	for i, array := range b {
		for j, cell := range array {
			if cell == 0 {
				m = append(m, Move{i, j, 0})
			}
		}

	}
	return m
}