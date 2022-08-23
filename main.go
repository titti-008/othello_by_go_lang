package main

import (
	"fmt"
	"github.com/gookit/color"
)

func main() {
	othello := NewOthello()

	for {
		gameOver := othello.CheckGame()
		color.Red.Print("P1 stone: ", othello.player1.stone, ", ")
		color.Green.Print("P2 stone: ", othello.player2.stone, "\n")
		if gameOver {
			fmt.Printf("GameOver! Winner is %s\n", othello.winner)
			break
		}

		fmt.Println("Please input integer x and y. exp) 3 2")
		fmt.Println("( If you skip this turn, input '8 8' )")

		if othello.turn {
			color.Red.Println("Now is the Player1 (○) turn.")
			color.Red.Print("input: ")
		} else {
			color.Green.Println("Now is the Player2 (x) turn.")
			color.Green.Print("input: ")
		}

		var x int
		var y int

		fmt.Scan(&x, &y)

		if x == 8 || y == 8 {
			othello.SkipTurn()
			continue
		}

		err := othello.Put(x, y)
		if err != nil {
			color.Errorln(err)
		}
	}
}
