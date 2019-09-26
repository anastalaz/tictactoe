package tictactoe

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// GetMove asks the human player to give the coordinates
// of his move through the console.
func (h Human) GetMove(b Board) (m Move) {
	fmt.Printf("Player %d make a move:\n", h.ID+1)
	fmt.Printf("x: ")
	fmt.Scanf("%d", &m[0])
	fmt.Printf("y: ")
	fmt.Scanf("%d", &m[1])
	m[2] = h.ID + 1

	return m
}

// GetMove just makes a random move.
func (r RandomAI) GetMove(b Board) (m Move) {
	rand.Seed(time.Now().UTC().UnixNano())

	moves := b.LegalMoves()
	m = moves[rand.Intn(len(moves))]
	m[2] = r.ID + 1
	return m
}

// GetMove uses the Min Max algorithm to make a move.
func (r MinMaxAI) GetMove(b Board) (m Move) {
	bestScore := 0

	moves := b.LegalMoves()
	for _, move := range moves {
		bc := b
		move[2] = r.ID + 1
		bc = bc.NewMove(move)
		opp := math.Abs(float64(r.ID - 1))
		score := minMaxScore(bc, int(opp), r.ID)
		if score >= bestScore {
			m = move
			bestScore = score
		}
	}
	return m
}

// minMaxScore calculates the score for a Position. If the Position
// is wining the function returns 10, if it is loosing -10 and if it
// is a draw 0.
func minMaxScore(board Board, playerToMove int, playerToOptimize int) int {
	winner := board.GetWinner()
	if winner != 0 {
		if winner == playerToOptimize+1 {
			return 10
		} else {
			return -10
		}

	} else if board.IsFull() {
		return 0
	}

	legalMoves := board.LegalMoves()

	var scores []int
	for _, move := range legalMoves {
		bc := board
		move[2] = playerToMove + 1
		bc = bc.NewMove(move)
		opp := math.Abs(float64(playerToMove - 1))

		oppBestResponseScore := minMaxScore(bc, int(opp), playerToOptimize)
		scores = append(scores, oppBestResponseScore)
	}
	sort.Ints(scores)

	if playerToMove == playerToOptimize {
		return scores[len(scores)-1]
	} else {
		return scores[0]
	}

}
