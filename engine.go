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

// Battle starts the game engine and lets 2 AIs play against each other.
func Battle(p1, p2 Player, battles int) {
	winsP1 := 0
	winsP2 := 0
	draws := 0

	for i := 0; i < battles; i++ {
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
			// board.Render()
			// fmt.Println("")
			winner := board.GetWinner()

			if winner != 0 {
				// fmt.Printf("The winner is Player %d\n", winner)

				if winner == 1 {
					winsP1++
				}
				winsP2++
				break
			}
			if board.IsFull() {
				// fmt.Println("The game is a draw")
				draws++
				break
			}
			round++
		}
	}

	fmt.Printf("Player 1 wins: %d (%.1f%%) - Player 2 wins: %d (%.1f%%) - Draws: %d\n", winsP1, float64(winsP1)*100/float64(battles), winsP2, float64(winsP2)*100/float64(battles), draws)

}
