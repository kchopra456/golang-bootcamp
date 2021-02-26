package main

func condStatement() {
	x := 5

	if x > 4 {
		println("x > 4")
	}

	if x > 0 && x < 10 {
		println("0 < x < 10")
	}

	if multiple := x * 2; multiple > 5 {
		println("x, 2 multiple > 5")
	}
}
