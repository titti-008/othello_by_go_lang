package main

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	player1 := player{}
	player1.New("Player1")

	if player1.name != "Player1" {
		t.Errorf("when initialize player, have to name Player1, but got %s", player1.name)
	}

	for y := range player1.board {
		for x := range player1.board {
			if player1.board[y][x] != false {
				t.Errorf("player1.board[%d][%d]の値がfalseではありません", y, x)
			}
		}
	}
}

func TestPrintPut(t *testing.T) {
	player := player{name: "Player1"}

	message := player.PrintPut(3, 2)
	if message != "Player1 put stone (x: 3, y: 2)" {
		t.Errorf("want `Player1 put stone (x: 3, y:2)` but got %s", message)
	}
}
