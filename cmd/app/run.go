package main

import (
	"fmt"
)

func main() {

	type figure struct {
		color  bool
		symbol string
		price  int
	}

	type board struct {
		cell [8][8]figure
	}

	var gameBoard board

	gameBoard.cell[0][1] = figure{true, "⛀", 0}

	for i := 0; i < len(gameBoard.cell); i++ {
		for j := 0; j < len(gameBoard.cell); j++ {
			if len(gameBoard.cell[i][j].symbol) == 0 {
				if (j+i)%2 == 0 {
					fmt.Print("■")
				} else {
					fmt.Print("□")
				}
			}
			print(gameBoard.cell[i][j].symbol)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}

}
