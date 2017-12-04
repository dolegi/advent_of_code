package main

import (
	"fmt"
	"math"
)

const (
	input = 325489
)

func main() {
	level := 0
	x := 0
	y := 0
	base := 2
	for i := 1; i <= input; i++ {
		switch {
		case i == base+(8*level):
			base = i
			level += 1
			x = level
			y = 1 - level
		case i < base+(level*2):
			y += 1
		case i < base+(level*4):
			x -= 1
		case i < base+(level*6):
			y -= 1
		case i < base+(level*8):
			x += 1
		}
	}
	steps := int(math.Abs(float64(x)) + math.Abs(float64(y)))
	fmt.Printf("%d %d | x: %d y: %d level: %d  base: %d\n", input, steps, x, y, level, base)
}
