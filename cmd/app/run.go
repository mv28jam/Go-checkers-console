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

	for i := 0; i < len(gameBoard.cell); i++ {
		println(gameBoard.cell[i][0].symbol)
	}

	println("ff")
}
