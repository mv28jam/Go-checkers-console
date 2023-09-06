package main

import (
	"fmt"
)

type figure struct {
	color  bool
	symbol string
	price  int
}

type board struct {
	cell [8][8]figure
}

func setFigures(gameBoard *board) {
	for i := 0; i < len(gameBoard.cell); i++ {
		for j := 0; j < 3; j++ {
			if (j+i)%2 != 0 {
				gameBoard.cell[j][i] = figure{true, "⛀", 1}
			}
		}
		for j := 5; j < 8; j++ {
			if (j+i)%2 != 0 {
				gameBoard.cell[j][i] = figure{false, "⛂", 1}
			}
		}
	}
}

func printBoard(gameBoard board) {
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

func main() {

	var gameBoard board

	setFigures(&gameBoard)
	printBoard(gameBoard)

}
