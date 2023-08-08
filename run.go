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

	println("ff")
}
