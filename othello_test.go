package main

import (
	"testing"
)




func TestNewOthello (t *testing.T) {

	othello := NewOthello()

	if othello.player1.name != "Player1" {
		t.Errorf("when initialize othello, have to name player1, but got %s", othello.player1.name)
	}


	if othello.player2.name != "Player2" {
		t.Errorf("when initialize othello, have to name player2, but got %s", othello.player2.name)
	}

	if othello.turn != true {
		t.Error("turn of after initialization is 'true', but got: ", othello.turn)
	}

	if othello.player1.board[4][3] != true ||  othello.player1.board[3][4] != true ||
	othello.player2.board[3][3] != true || othello.player2.board[4][4] != true {
		t.Error("Initialize bourd is put 4 stone, but not", othello.PrintBoard())
	}
}


func TestIsEmpty(t *testing.T) {
	othello := NewOthello()

	othello.player1.board[1][5] = true

	isEmpty := othello.IsEmpty(5, 1)
	if isEmpty != false {
		t.Errorf("IsEmpty(5, 1) ==  got %t; want false", isEmpty)
	}
}

func TestPut(t *testing.T) {
	othello := NewOthello()
	othello.player1.board[5][1] = true
	
	othello.turn = true

	err := othello.Put(1, 5)
	if err == nil {
		t.Error("shoud output error when put fillid cell, but does not raise error.", err)
	}

	err = othello.Put(3, 2)
	if err != nil {
		t.Error("shoud not output error when put empty cell, but it raise error.", err)
	}

	if othello.player1.board[2][3] != true {
		t.Errorf("if turn is true, player1 put stone, but result is %v.", othello.player1.board[2][3])
	}


	if othello.turn != false {
		t.Errorf("After player1 put stone, turn is expect 'false' but got '%t' ", othello.turn)
	}
}

func TestChangeTurn (t *testing.T) {
	othello := NewOthello()

	othello.turn = true

	othello.ChangeTurn()

	if othello.turn != false {
		t.Errorf("turn is expected change, but not change (othello.ChangeTurn)  true → %t", othello.turn)
	}
}


func TestPrintBoard (t *testing.T) {
	othello := NewOthello()
	
	board := `
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][x][o][ ][ ][ ]
[ ][ ][ ][o][x][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
`

	if othello.PrintBoard() != board {
		t.Error("expect:", board, "but got:", othello.PrintBoard())
	}

		
	board = `
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][o][x][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][x][o][ ][ ][ ]
[ ][ ][ ][o][x][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
[ ][ ][ ][ ][ ][ ][ ][ ]
`

	othello.Put(2, 1)
	othello.Put(3, 1)
	if othello.PrintBoard() != board {
		t.Error("expect:", board, "but got:", othello.PrintBoard())
	}
}