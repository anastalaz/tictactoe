package tictactoe

import "testing"

func TestGetWinner(t *testing.T) {

	t.Run("Horizontal player 1 win", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			var board Board
			board[0][i] = 1
			board[1][i] = 1
			board[2][i] = 1

			got := board.GetWinner()
			want := 1

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})

	t.Run("Vertical player 2 win", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			var board Board
			board[i][0] = 2
			board[i][1] = 2
			board[i][2] = 2

			got := board.GetWinner()
			want := 2

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})

	t.Run("Diagonal player 1 win", func(t *testing.T) {
		var board Board
		board[0][0] = 1
		board[1][1] = 1
		board[2][2] = 1

		got := board.GetWinner()
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("Diagonal player 2 win", func(t *testing.T) {
		var board Board
		board[2][0] = 2
		board[1][1] = 2
		board[0][2] = 2

		got := board.GetWinner()
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Draw", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			var board Board
			board[i][0] = 1
			board[i][1] = 2
			board[i][2] = 1

			got := board.GetWinner()
			want := 0

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})
}

func TestNewMove(t *testing.T) {
	var board Board
	board = board.NewMove(Move{0, 0, 1})

	got := board.NewMove(Move{0, 1, 2})
	want := Board{
		{1, 2, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestIsFull(t *testing.T) {
	var board Board
	board = Board{
		{1, 2, 1},
		{2, 1, 2},
		{2, 0, 2},
	}

	got := board.IsFull()
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}

	board = Board{
		{1, 2, 2},
		{2, 1, 1},
		{1, 2, 2},
	}

	got = board.IsFull()
	want = true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestLegalMoves(t *testing.T) {
	var board Board
	board = Board{
		{0, 2, 1},
		{2, 0, 2},
		{2, 1, 2},
	}

	got := board.LegalMoves()
	want := Move{0, 0, 0}
	want1 := Move{1, 1, 0}
	length := len(got)

	if got[0] != want || got[1] != want1 || length != 2 {
		t.Errorf("got %v want %v and %v", got, want, want1)
	}
}
