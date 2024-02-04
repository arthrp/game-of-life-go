package main

import (
	"fmt"
	"time"
)

const Width = 20
const Height = 20

type GameState struct {
	Field [Width][Height]bool
}

func (g *GameState) printState() {
	s := ""
	for y := 0; y < Width; y++ {
		for x := 0; x < Height; x++ {
			if g.Field[x][y] {
				s += "■ "
			} else {
				s += "□ "
			}
		}
		s += "\n"
	}

	fmt.Println(s)
}

func (g *GameState) initState() {
	g.Field[2][2] = true
	g.Field[2][3] = true
	g.Field[2][4] = true
}

func (g *GameState) countNeighbours(cellX int, cellY int) int {
	var count = 0

	for dy := -1; dy <= 1; dy++ {
		var y = (cellY + dy + Height) % Height
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			var x = (cellX + dx + Width) % Width
			if g.Field[x][y] {
				count += 1
			}
		}
	}

	return count
}

func (g *GameState) makeTurn() {
	var newField [Width][Height]bool

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			var n = g.countNeighbours(x, y)
			var cell = g.Field[x][y]
			if (cell && (n == 2 || n == 3)) || (!cell && n == 3) {
				newField[x][y] = true
			} else {
				newField[x][y] = false
			}
		}
	}

	g.Field = newField
}

func main() {
	var game GameState

	game.initState()
	for {
		//Clear screen
		fmt.Print("\x1B[2J")
		game.makeTurn()
		game.printState()
		time.Sleep(500 * time.Millisecond)
	}
}
