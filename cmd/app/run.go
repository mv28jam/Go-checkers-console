package main

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

	gameBoard.cell[0][0] = figure{true, "⛀", 0}

	for i := 0; i < len(gameBoard.cell); i++ {
		println(gameBoard.cell[i][0].symbol)
	}

}
