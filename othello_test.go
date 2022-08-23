package main

import (
	"fmt"
	"testing"
)

func TestNewOthello(t *testing.T) {

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

	if othello.player1.board[4][3] != true || othello.player1.board[3][4] != true ||
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

func TestCheckSurroundings(t *testing.T) {
	othello := NewOthello()

	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][?][x][o][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	lineList := othello.CheckSurroundings(2, 3)
	if len(lineList) != 1 {
		t.Error("when checking cell(2, 3) that have enemy cell in front of that, result want 1 lineList , but got: ", len(lineList), othello.PrintBoard())
	}

	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][?][ ][ ][ ][ ][ ]
	// [ ][ ][ ][x][o][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	lineList = othello.CheckSurroundings(2, 2)
	if len(lineList) != 0 {
		t.Error("when checking cell(2, 2) that have enemy cell in front of that, result want 1 lineList , but got: ", lineList, othello.PrintBoard())
	}

	// [?][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][x][o][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	lineList = othello.CheckSurroundings(0, 0)
	if len(lineList) != 0 {
		t.Error("when checking cell(0, 0) that have enemy cell in front of that, result want 0 lineList , but got: ", lineList, othello.PrintBoard())
	}

	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][?][ ][ ][ ][ ][ ]
	// [ ][ ][x][x][o][ ][ ][ ]
	// [ ][ ][o][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][o][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	othello.player1.board[5][5] = true
	othello.player1.board[4][2] = true
	othello.player2.board[3][2] = true
	lineList = othello.CheckSurroundings(2, 2)
	if len(lineList) != 2 {
		t.Error("when checking cell(2, 2) that have two enemy cell in front of that, result want 2 lineList , but got: ", len(lineList), othello.PrintBoard())
	}
	if len(lineList[0]) != 1 {
		t.Error("when checking cell(2, 2) that have two enemy cell in front of that, linelist[0] shoud have one cell , but got: ",
			len(lineList[0]), othello.PrintBoard())
	}

	if len(lineList[1]) != 2 {
		t.Error("when checking cell(2, 2) that have two enemy cell in front of that, linelist[1] shoud have two cell , but got: ",
			len(lineList[1]), othello.PrintBoard())
	}
}

func CheckZeroEnemyLine(t *testing.T, lineList []pos) {
	if len(lineList) != 0 {
		t.Error("when not enemy stone in CheckLine, lineList length is 0, but got:", len(lineList))
	}
}

func ChecklineList(t *testing.T, lineList, expectPosLiest []pos, liLeng int) {
	if lineList[0] != expectPosLiest[0] {
		t.Error("when is there enemy in check next line, lineList[0] is that position, but got:", lineList[0])
	}
	if len(lineList) != liLeng {
		t.Error("when are there enemy consecutive in CheckLine, lineList length is 1, but got:", len(lineList))
	}

}

// 重複したテストをDRY

func TestPut(t *testing.T) {
	othello := NewOthello()
	othello.player1.board[5][1] = true

	othello.turn = true

	err := othello.Put(-1, 11)
	if err == nil {
		t.Error("shoud output error when put outside board, but does not raise error.", err)
	}

	err = othello.Put(0, 0)
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

func TestChangeTurn(t *testing.T) {
	othello := NewOthello()

	othello.turn = true

	othello.ChangeTurn()

	if othello.turn != false {
		t.Errorf("turn is expected change, but not change (othello.ChangeTurn)  true → %t", othello.turn)
	}
}

func TestPrintBoard(t *testing.T) {
	othello := NewOthello()

	board := `
   0  1  2  3  4  5  6  7
0 [ ][ ][ ][ ][ ][ ][ ][ ]
1 [ ][ ][ ][ ][ ][ ][ ][ ]
2 [ ][ ][ ][ ][ ][ ][ ][ ]
3 [ ][ ][ ][<green>x</>][<red>o</>][ ][ ][ ]
4 [ ][ ][ ][<red>o</>][<green>x</>][ ][ ][ ]
5 [ ][ ][ ][ ][ ][ ][ ][ ]
6 [ ][ ][ ][ ][ ][ ][ ][ ]
7 [ ][ ][ ][ ][ ][ ][ ][ ]
`

	if othello.PrintBoard() != board {
		t.Error("expect:", board, "but got:", othello.PrintBoard())
	}

	board = `
   0  1  2  3  4  5  6  7
0 [ ][ ][ ][ ][ ][ ][ ][ ]
1 [ ][ ][ ][ ][ ][ ][ ][ ]
2 [ ][ ][ ][ ][<green>x</>][ ][ ][ ]
3 [ ][ ][<red>o</>][<red>o</>][<green>x</>][ ][ ][ ]
4 [ ][ ][ ][<red>o</>][<green>x</>][ ][ ][ ]
5 [ ][ ][ ][ ][ ][ ][ ][ ]
6 [ ][ ][ ][ ][ ][ ][ ][ ]
7 [ ][ ][ ][ ][ ][ ][ ][ ]
`

	othello.Put(2, 3)
	othello.Put(4, 2)
	if othello.PrintBoard() != board {
		t.Error("expect:", board, "but got:", othello.PrintBoard())
	}
}

func TestCheckLine(t *testing.T) {
	othello := NewOthello()

	// Player1 turn
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][↓][ ][ ][ ]
	// [ ][ ][ ][x][o][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	var position = pos{x: 2, y: 3}
	var delta = pos{x: 1, y: 0}

	lineList := othello.CheckLine(position, delta)
	var expectPosLiest []pos = []pos{{x: 3, y: 3}}
	ChecklineList(t, lineList, expectPosLiest, 1)

	othello.Put(3, 2) // put by Player1
	// Player2 turn
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][o][↓][ ][ ][ ]
	// [ ][ ][ ][o][o][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	position = pos{x: 4, y: 2}
	delta = pos{x: 0, y: 1}
	lineList = othello.CheckLine(position, delta)
	expectPosLiest = []pos{{x: 4, y: 3}}
	ChecklineList(t, lineList, expectPosLiest, 1)

	othello.Put(4, 2) // put by Player2
	// Player1 turn
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][o][x][└][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]

	fmt.Println(othello.PrintBoard())
	position = pos{x: 5, y: 2}
	delta = pos{x: -1, y: 1}
	lineList = othello.CheckLine(position, delta)
	expectPosLiest = []pos{{x: 4, y: 3}}
	ChecklineList(t, lineList, expectPosLiest, 1)

	// [┘][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]

	// Player1 turn
	position = pos{x: 0, y: 0}
	delta = pos{x: 1, y: 1}
	lineList = othello.CheckLine(position, delta)
	CheckZeroEnemyLine(t, lineList)

	// When there is my stone in front of me
	position = pos{x: 4, y: 2}
	delta = pos{x: 0, y: 1}
	lineList = othello.CheckLine(position, delta)
	CheckZeroEnemyLine(t, lineList)

	// When there is empty cell in front of me
	delta = pos{x: 0, y: -1}
	lineList = othello.CheckLine(position, delta)
	CheckZeroEnemyLine(t, lineList)

	// When there is wall in front of me
	position = pos{x: 0, y: 0}
	delta = pos{x: 0, y: -1}
	lineList = othello.CheckLine(position, delta)
	CheckZeroEnemyLine(t, lineList)

	// When there is wall in front of me
	position = pos{x: 7, y: 7}
	delta = pos{x: 1, y: 1}
	lineList = othello.CheckLine(position, delta)
	CheckZeroEnemyLine(t, lineList)
}

func TestCheckCanPut(t *testing.T) {
	othello := NewOthello()

	// When checked filled cell.
	result := othello.ChackCanPut(3, 4)
	if result != false {
		t.Error("when check cell that filled cell, result want 'false', but ChackCanPut result is: ", result)
	}

	// When checked cell with empty surroundings.
	result = othello.ChackCanPut(1, 1)
	if result != false {
		t.Error("when check cell with empty surroundings, result want 'false', but ChackCanPut result is: ", result)
	}
}

func TestTurnOverArroundCell(t *testing.T) {
	othello := NewOthello()

	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][!][ ][ ][ ][ ]
	// [ ][ ][ ][x][o][ ][ ][ ]
	// [ ][ ][ ][o][x][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	// [ ][ ][ ][ ][ ][ ][ ][ ]
	surroundings := othello.CheckSurroundings(3, 2)
	othello.TurnOverArroundCells(surroundings)

	if othello.player1.board[3][3] != true || othello.player2.board[3][3] != false {
		t.Error("when do TurnOverArroundCell(3,2), (3,3) should be change player1's stone, but that is not.",
			"othello.player1.board[3][3]:", othello.player1.board[3][3],
			"othello.player2.board[3][3]", othello.player2.board[3][3],
		)
	}
}

func TestTurnOver(t *testing.T) {
	othello := NewOthello()
	if othello.player1.board[3][3] != false || othello.player2.board[3][3] != true {
		t.Error("when before TurnOver(3,3) this cell should be player2's stone, but that is not.",
			"othello.player1.board[3][3]:", othello.player1.board[3][3],
			"othello.player2.board[3][3]", othello.player2.board[3][3],
		)
	}
	othello.TurnOver(3, 3)
	if othello.player1.board[3][3] != true || othello.player2.board[3][3] != false {
		t.Error("when do TurnOver(3,3) should be change player1's stone, but that is not.",
			"othello.player1.board[3][3]:", othello.player1.board[3][3],
			"othello.player2.board[3][3]", othello.player2.board[3][3],
		)
	}
}

func TestSkipTurn(t *testing.T) {

	othello := NewOthello()
	if othello.turn != true {
		t.Error("turn have to true after initialize othello, but got:", othello.turn)
	}
	othello.SkipTurn()
	if othello.turn != false {
		t.Error("turn have to changing false from true after SlipTurn(), but got:", othello.turn)
	}

}

func TestCountStone(t *testing.T) {
	othello := NewOthello()

	othello.CountStone()

	if othello.player1.stone != 2 {
		t.Error("When after initialize, player1 stone count is 2,but got:", othello.player1.stone)
	}

	if othello.player2.stone != 2 {
		t.Error("When after initialize, player2 stone count is 2,but got:", othello.player2.stone)
	}

	othello.Put(3, 2)
	othello.CountStone()
	if othello.player1.stone != 4 {
		t.Error("Player1 put stone, player1 expect stone count 2->4, but got :", othello.player1.stone)
	}

	if othello.player2.stone != 1 {
		t.Error("Pleayer1 put stone, plaper2 expect stone count 2->1, but got :", othello.player2.stone)
	}
}

func TestIsGameOver(t *testing.T) {
	othello := NewOthello()

	othello.player1.board[4][3] = false
	othello.player2.board[4][3] = true

	othello.player1.board[3][4] = false
	othello.player2.board[3][4] = true

	//    0  1  2  3  4  5  6  7
	// 0 [ ][ ][ ][ ][ ][ ][ ][ ]
	// 1 [ ][ ][ ][ ][ ][ ][ ][ ]
	// 2 [ ][ ][ ][ ][ ][ ][ ][ ]
	// 3 [ ][ ][ ][x][x][ ][ ][ ]
	// 4 [ ][ ][ ][x][x][ ][ ][ ]
	// 5 [ ][ ][ ][ ][ ][ ][ ][ ]
	// 6 [ ][ ][ ][ ][ ][ ][ ][ ]
	// 7 [ ][ ][ ][ ][ ][ ][ ][ ]

	// when Player1 is not able to put
	gameOver := othello.CheckGame()

	if gameOver != true {
		t.Error("when there is any empty cell and player have no cell that is able to put, othello.CheckGame() to be GameOver, but got:",
			gameOver, othello.PrintBoard())
	}

	if othello.winner != "Player2" {
		t.Error("want: othello.winner == 'draw',  but got:", othello.winner)
	}

	for y := range othello.player1.board {
		for x := range othello.player1.board[y] {
			if (x+y)%2 == 0 {
				othello.player1.board[y][x] = true
				othello.player2.board[y][x] = false
			} else {
				othello.player2.board[y][x] = true
				othello.player1.board[y][x] = false
			}
		}
	}

	gameOver = othello.CheckGame()

	if gameOver != true {
		t.Error("when stone count is 64, it's game over, and return true, but got:", gameOver, othello.player1.stone, othello.player2.stone)
	}

	if othello.winner != "draw" {
		t.Error("want: othello.winner == 'draw',  but got:", othello.winner)
	}

	othello.player1.board[0][0] = false
	othello.player2.board[0][0] = true

	gameOver = othello.CheckGame()
	if othello.winner != "Player2" {
		t.Error("want: othello.winner == 'Player2',  but got:", othello.winner)
	}
}

func TestCheckEnableCells(t *testing.T) {
	othello := NewOthello()

	enableCells := othello.CheckEnableCells()

	firstCell := pos{x: 3, y: 2}
	if enableCells[0] != firstCell {
		t.Error("want othello.CheckEnableCells()[0] == pos(x:3, y:2), but enableCells got:", enableCells)
	}

	if len(enableCells) != 4 {
		t.Error("want len(enableCells) == 4, but got:", len(enableCells))
	}
}
