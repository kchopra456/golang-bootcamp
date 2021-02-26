package main

func fizzbuzz() {
	for x := 1; x <= 20; x++ {
		switch {
		case x%15 == 0:
			println("fizz buzz")
		case x%3 == 0:
			println("fizz")

		case x%5 == 0:
			println("buzz")

		default:
			println(x)
		}
	}
}
