
// CheckArround
// EnabledTurnStone
// SkipTurn
// Enabled

// TurnOverStone




package main

import (
	"fmt"
)

type othello struct {
	player1 player
	player2 player
	turn 	bool
}

// オセロゲームの初期化
func NewOthello () othello{
	o := othello{turn: true}
	
	o.player1.New("Player1")
	o.player2.New("Player2")

	o.player1.board[4][3] = true 
	o.player1.board[3][4] = true 
	o.player2.board[3][3] = true 
	o.player2.board[4][4] = true

	fmt.Println("Start Othello!!")

	fmt.Printf(o.PrintBoard())

	return o
}

// 指定したセルに石を置けるか確かめる
func (o *othello) IsEmpty (x, y uint) bool{
	p1 := o.player1.board[y][x]
	p2 := o.player2.board[y][x]

	return !p1 && !p2
}

func (o *othello) Put (x, y uint) error{

	if !o.IsEmpty(x, y) {
		return fmt.Errorf("This cell is not empty.")
	}
	if o.turn {
		fmt.Print(o.player1.PrintPut(x, y))
		o.player1.board[y][x] = true
	} else if !o.turn {
		fmt.Print(o.player2.PrintPut(x, y))
		o.player2.board[y][x] = true
	}

	fmt.Printf(o.PrintBoard())

	o.ChangeTurn()
	return nil
}

func (o *othello) ChangeTurn () {
	o.turn = !o.turn
}

func (o *othello) PrintBoard () (board string) {
	board = "\n"
	for y := range o.player1.board {
		for x := range o.player1.board[y] {
			stone := " "
			if o.player1.board[y][x] {
				stone = "o"
			} else if o.player2.board[y][x] {
				stone = "x"
			}
			board += fmt.Sprintf("[%s]", stone)
		}
		board += "\n"
	}

	return
}