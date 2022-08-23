package main

import (
	"fmt"
	"github.com/gookit/color"
)

type othello struct {
	player1 player
	player2 player
	turn    bool
	winner  string
}

type pos struct {
	x int
	y int
}

type lineList []pos

// Inisialize game
func NewOthello() othello {
	o := othello{turn: true}

	o.player1.New("Player1")
	o.player2.New("Player2")

	o.player1.board[4][3] = true
	o.player1.board[3][4] = true
	o.player2.board[3][3] = true
	o.player2.board[4][4] = true

	fmt.Println("Start Othello!!")

	color.Print(o.PrintBoard())

	return o
}

func (o *othello) SkipTurn() {
	color.Print(o.PrintBoard())
	p := ""
	if o.turn {
		p = o.player1.name
	} else {
		p = o.player2.name
	}
	fmt.Printf("%s skip turn!\n", p)
	o.ChangeTurn()

}

func (o *othello) Put(x, y int) error {
	if x < 0 || x > 7 || y < 0 || y > 7 {
		color.Print(o.PrintBoard())
		return fmt.Errorf("x and y is only 0 to 7. but you input (x: %d, x: %d).", x, y)
	}
	if !o.ChackCanPut(x, y) {
		color.Print(o.PrintBoard())
		return fmt.Errorf("you can't put cell(x: %d, y: %d).", x, y)
	}
	surroundings := o.CheckSurroundings(x, y)
	o.TurnOverArroundCells(surroundings)
	if o.turn {
		fmt.Print(o.player1.PrintPut(x, y))
		o.player1.board[y][x] = true
	} else if !o.turn {
		fmt.Print(o.player2.PrintPut(x, y))
		o.player2.board[y][x] = true
	}
	color.Print(o.PrintBoard())
	o.ChangeTurn()
	return nil
}

func (o *othello) ChackCanPut(x, y int) bool {
	if !o.IsEmpty(x, y) {
		return false
	}

	surroundings := o.CheckSurroundings(int(x), int(y))
	if len(surroundings) == 0 {
		return false
	}

	return true
}

// Check if the specified cell is empty.
func (o *othello) IsEmpty(x, y int) bool {
	p1 := o.player1.board[y][x]
	p2 := o.player2.board[y][x]

	return !p1 && !p2
}

func (o *othello) CheckSurroundings(x, y int) (surroundings []lineList) {
	for dy := -1; dy <= 1; dy++ {
		if y+dy < 0 || y+dy > 7 {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			if x+dx < 0 || x+dx > 7 {
				continue
			}
			if dx == 0 && dy == 0 {
				continue
			}
			lineList := o.CheckLine(pos{x: x, y: y}, pos{x: dx, y: dy})
			if len(lineList) == 0 {
				continue
			}
			surroundings = append(surroundings, lineList)
		}
	}

	return surroundings
}

func (o *othello) CheckLine(position, delta pos) (lineList lineList) {
	for i := 1; i < 7; i++ {
		lineX := position.x + (delta.x * i)
		lineY := position.y + (delta.y * i)

		if lineX < 0 || lineX > 7 || lineY < 0 || lineY > 7 {
			// when there is wall in front of me, return  empty list
			return lineList[:0]
		}

		if o.IsEmpty(lineX, lineY) {
			// when there is empty cell in front of me, return  empty list
			return lineList[:0]
		}

		// Player1's turn if o.turn is true.
		if o.turn {
			if o.player1.board[lineY][lineX] == true {
				return lineList
			}
			if o.player2.board[lineY][lineX] == true {
				lineList = append(lineList, pos{x: lineX, y: lineY})
			}
		} else {
			if o.player2.board[lineY][lineX] == true {
				return lineList
			}
			if o.player1.board[lineY][lineX] == true {
				lineList = append(lineList, pos{x: lineX, y: lineY})
			}
		}
	}
	return lineList
}

func (o *othello) ChangeTurn() {
	o.turn = !o.turn
}

// return board state strings
// Player1 -> o
// Player2 -> x
// `
//    0  1  2  3  4  5  6  7
// 0 [ ][ ][ ][ ][ ][ ][ ][ ]
// 1 [ ][ ][ ][ ][ ][ ][ ][ ]
// 2 [ ][ ][ ][ ][ ][ ][ ][ ]
// 3 [ ][ ][ ][x][o][ ][ ][ ]
// 4 [ ][ ][ ][o][x][ ][ ][ ]
// 5 [ ][ ][ ][ ][ ][ ][ ][ ]
// 6 [ ][ ][ ][ ][ ][ ][ ][ ]
// 7 [ ][ ][ ][ ][ ][ ][ ][ ]
// .
func (o *othello) PrintBoard() (board string) {
	board = "\n   0  1  2  3  4  5  6  7\n"
	for y := range o.player1.board {
		board += fmt.Sprintf("%d ", y)
		for x := range o.player1.board[y] {
			stone := " "
			if o.player1.board[y][x] {
				stone = "<red>o</>"
			} else if o.player2.board[y][x] {
				stone = "<green>x</>"
			}
			board += fmt.Sprintf("[%s]", stone)
		}
		board += "\n"
	}
	return
}

func (o *othello) TurnOverArroundCells(surroundings []lineList) {
	for _, lineList := range surroundings {
		for _, enemyPos := range lineList {
			o.TurnOver(enemyPos.x, enemyPos.y)
		}
	}
}

func (o *othello) TurnOver(x, y int) {
	o.player1.board[y][x] = !o.player1.board[y][x]
	o.player2.board[y][x] = !o.player2.board[y][x]
}

func (o *othello) CountStone() (totalStone int) {
	o.player1.stone = 0
	o.player2.stone = 0
	for y := range o.player1.board {
		for x := range o.player2.board[y] {
			if o.player1.board[y][x] == true {
				o.player1.stone++
			}
			if o.player2.board[y][x] == true {
				o.player2.stone++
			}
		}
	}
	totalStone = o.player1.stone + o.player2.stone
	return totalStone
}

func (o *othello) CheckGame() (gameOver bool) {
	gameOver = false

	totalStone := o.CountStone()

	if totalStone >= 64 {
		gameOver = true
	}

	o.CheckEnableCells()
	if o.player1.enablePos == 0 && o.player2.enablePos == 0 {
		gameOver = true
	}

	o.CheckWinner()
	return gameOver
}

func (o *othello) CheckEnableCells() (cells []pos) {
	for y := range o.player1.board {
		for x := range o.player1.board[y] {
			if o.ChackCanPut(x, y) {
				cells = append(cells, pos{x: x, y: y})
			}
		}
	}
	if o.turn {
		o.player1.enablePos = len(cells)
	} else {
		o.player2.enablePos = len(cells)
	}
	return cells
}

func (o *othello) CheckWinner() {
	if o.player1.stone > o.player2.stone {
		o.winner = o.player1.name
	} else if o.player1.stone < o.player2.stone {
		o.winner = o.player2.name
	} else if o.player1.stone == o.player2.stone {
		o.winner = "draw"
	}
}
