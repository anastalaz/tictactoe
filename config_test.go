package tictactoe

import "testing"

func TestCreatePlayer(t *testing.T) {
	want1 := Human{ID: 0}
	got1 := createPlayer("human", 0)

	if got1 != want1 {
		t.Errorf("got %d want %d", got1, want1)
	}

	want2 := RandomAI{ID: 1}
	got2 := createPlayer("randomai", 1)

	if got2 != want2 {
		t.Errorf("got %d want %d", got2, want2)
	}
}
