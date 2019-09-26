package tictactoe

import (
	"errors"
	"fmt"
	"reflect"
)

// Player is the interface that the various agents have
// to implement.
type Player interface {
	GetMove(b Board) Move
}

// Human is the human agent with ID 1 if it is player 1
// or ID 2 if it is player 2.
type Human struct {
	ID int
}

// RandomAI is the randomai agent with ID 1 if it is player 1
// or ID 2 if it is player 2. The Agent just makes a random
// move in one of the free squares.
type RandomAI struct {
	ID int
}

// MinMaxAI is the minmaxai agent with ID 1 if it is player 1
// or ID 2 if it is player 2. The Agent uses the Min Max algorithm
// to make a move.
type MinMaxAI struct {
	ID int
}

// Config needs two inputs from the user, the agents for
// player one and two. The function creates and returns
// the two players or an error if the wrong input was given.
func Config() (p1, p2 Player, err error) {
	var a int

	fmt.Printf("Player 1: Human (1), RandomAI (2) or MinMaxAI (3): ")
	fmt.Scanf("%d", &a)
	if a == 2 {
		p1 = createPlayer("randomai", 0)
	} else if a == 1 {
		p1 = createPlayer("human", 0)
	} else if a == 3 {
		p1 = createPlayer("minmaxai", 0)
	} else {
		return nil, nil, errors.New("Wrong input for player 1")
	}

	fmt.Printf("Player 2: Human (1), RandomAI (2) or MinMaxAI (3): ")
	fmt.Scanf("%d", &a)
	if a == 2 {
		p2 = createPlayer("randomai", 1)
	} else if a == 1 {
		p2 = createPlayer("human", 1)
	} else if a == 3 {
		p2 = createPlayer("minmaxai", 1)
	} else {
		return nil, nil, errors.New("Wrong input for player 2")
	}
	return p1, p2, nil
}

func createPlayer(action string, ID int64) Player {
	var h Human
	var r RandomAI
	var m MinMaxAI
	if action == "human" {
		hType := reflect.TypeOf(h)
		pl1 := reflect.New(hType)
		pl1.Elem().Field(0).SetInt(ID)
		p1 := pl1.Elem().Interface().(Human)
		return p1
	} else if action == "randomai" {
		rType := reflect.TypeOf(r)
		pl2 := reflect.New(rType)
		pl2.Elem().Field(0).SetInt(ID)
		p2 := pl2.Elem().Interface().(RandomAI)
		return p2
	}
	mType := reflect.TypeOf(m)
	pl3 := reflect.New(mType)
	pl3.Elem().Field(0).SetInt(ID)
	p3 := pl3.Elem().Interface().(MinMaxAI)
	return p3

}
