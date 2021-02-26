package main

import "fmt"

type point struct {
	x int
	y int
}

func (p *point) move(x int, y int) {
	p.x += x
	p.y += y
}

func introStructs() {
	p := point{1, 2}
	p.move(2, 2)
	fmt.Printf("%+v", p)
}
