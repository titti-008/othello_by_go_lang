package main

import (
	"fmt"
)

type player struct {
	name      string
	board     [8][8]bool
	stone     int
	enablePos int
}

func (p *player) PrintPut(x, y int) string {
	return fmt.Sprintf("%s put stone (x: %d, y: %d)", p.name, x, y)
}

func (p *player) New(name string) {
	p.name = name
	for y := range p.board {
		for x := range p.board {
			p.board[y][x] = false
		}
	}
}
