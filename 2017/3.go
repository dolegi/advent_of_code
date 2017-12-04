package main

import (
	"fmt"
	"math"
)

const (
	input = 325489
)

type square struct {
	x     int
	y     int
	value int
}

func main() {
	part2()
}

func nextTo(x1 int, y1 int, x2 int, y2 int) bool {
	return (x1 >= (x2 - 1) && x1 <= (x2 + 1)) && (y1 >= (y2 - 1) && y1 <= (y2 + 1))
}

func part2() {
	var prevLevel []square
	var currentLevel []square
	level := 0
	x := 0
	y := 0
	base := 2
	currentLevel = append(currentLevel, square{x, y, 1})
	for i := 2; i <= input; i++ {
		switch {
		case i == base+(8*level):
			prevLevel = currentLevel
			currentLevel = []square{}
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

		value := 0

		for _, s := range prevLevel {
			if nextTo(x, y, s.x, s.y) {
				value += s.value
			}
		}
		for _, s := range currentLevel {
			if nextTo(x, y, s.x, s.y) {
				value += s.value
			}
		}
		currentLevel = append(currentLevel, square{x, y, value})
		if value > input {
			break
		}
	}
	fmt.Println(currentLevel[len(currentLevel) -1].value)
}

func part1() {
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
