// Package tictactoe provides various AI agents for the Tic Tac Toe game.
package tictactoe

import "fmt"

// Play starts the game engine and needs 2 players as input.
// The engine ends when there is a winner or a draw.
func Play(p1, p2 Player) {
	board := NewBoard()
	var currentPlayer Player
	round := 0

	for {
		if round%2 == 0 {
			currentPlayer = p1
		} else {
			currentPlayer = p2
		}

		m := currentPlayer.GetMove(board)
		board = board.NewMove(m)
		board.Render()
		fmt.Println("")
		winner := board.GetWinner()

		if winner != 0 {
			fmt.Printf("The winner is Player %d\n", winner)
			break
		}
		if board.IsFull() {
			fmt.Println("The game is a draw")
			break
		}
		round++
	}
}
