package main

import (
	"fmt"
	"strings"
)

const (
	input = `...#.#.####.....#.##..###
##.#.###..#.....#.##...#.
..#.##..#.#.##.#...#..###
###...##....###.#..#...#.
...#..#.........##..###..
#..#.#.#.#.#.#.#.##.####.
#...#.##...###...##..#..#
##...#.###..###...####.##
###..#.#####.##..###.#.##
#..#....#.##..####...####
...#.#......###.#..#..##.
.#.#...##.#.#####..###.#.
.....#..##..##..###....##
#.#..###.##.##.#####.##..
###..#..###.##.#..#.##.##
.#######.###....######.##
..#.#.###.##.##...###.#..
#..#.####...###..###..###
#...#..###.##..##...#.#..
........###..#.#.##..##..
.#############.#.###..###
##..#.###....#.#..#..##.#
..#.#.#####....#..#####..
.#.#..#...#...##.#..#....
##.#..#..##........#..##.`

	testInput = `..#
#..
...`
)

type direction struct {
	x int
	y int
}

type turtle struct {
	x int
	y int
	d string
}

func printGrid(grid [][]byte) {
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'c':
				fmt.Printf("%s", ".")
			case 'w':
				fmt.Printf("%s", "w")
			case 'f':
				fmt.Printf("%s", "f")
			case 'i':
				fmt.Printf("%s", "#")
			default:
				fmt.Printf("%s", ".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func move(t turtle) (int, int) {
	switch t.d {
	case "left":
		return 0, -1
	case "right":
		return 0, 1
	case "up":
		return -1, 0
	case "down":
		return 1, 0
	default:
		return 0, 0
	}
}

func newDirection(d string, n string) string {
	switch d {
	case "left":
		switch n {
		case "left":
			return "down"
		case "right":
			return "up"
		}
	case "right":
		switch n {
		case "left":
			return "up"
		case "right":
			return "down"
		}
	case "up":
		switch n {
		case "left":
			return "left"
		case "right":
			return "right"
		}
	case "down":
		switch n {
		case "left":
			return "right"
		case "right":
			return "left"
		}
	}
	return ""
}

func reverseDirection(d string) string {
	switch d {
	case "left":
		return "right"
	case "right":
		return "left"
	case "up":
		return "down"
	case "down":
		return "up"
	}
	return ""
}

func main() {
	grid := [][]byte{}
	for i, l := range strings.Split(input, "\n") {
		grid = append(grid, []byte{})
		for _, c := range strings.Split(l, "") {
			if c == "#" {
				grid[i] = append(grid[i], 'i')
			} else {
				grid[i] = append(grid[i], 'c')
			}
		}
	}

	infectedCount := 0
	width := (len(grid) - 1) / 2
	t := turtle{width, width, "up"}
	for i := 0; i < 10000000; i++ {
		switch grid[t.x][t.y] {
		case 'i':
			t.d = newDirection(t.d, "right")
			x, y := move(t)
			grid[t.x][t.y] = 'f'
			t.x = t.x + x
			t.y = t.y + y
		case 'c':
			t.d = newDirection(t.d, "left")
			x, y := move(t)
			grid[t.x][t.y] = 'w'
			t.x = t.x + x
			t.y = t.y + y
		case 'w':
			grid[t.x][t.y] = 'i'
			infectedCount++
			x, y := move(t)
			t.x = t.x + x
			t.y = t.y + y
		case 'f':
			t.d = reverseDirection(t.d)
			x, y := move(t)
			grid[t.x][t.y] = 'c'
			t.x = t.x + x
			t.y = t.y + y
		}

		if t.x < 0 {
			grid = append([][]byte{[]byte{}}, grid...)
			for _ = range grid[1] {
				grid[0] = append(grid[0], 'c')
			}
			t.x = t.x + 1
		} else if t.x >= len(grid) {
			grid = append(grid, []byte{})
			for _ = range grid[1] {
				grid[len(grid)-1] = append(grid[len(grid)-1], 'c')
			}
		}

		if t.y < 0 {
			for i := range grid {
				grid[i] = append([]byte{'c'}, grid[i]...)
			}
			t.y = t.y + 1
		} else if t.y >= len(grid[0]) {
			for i := range grid {
				grid[i] = append(grid[i], 'c')
			}
		}
	}
	fmt.Println(infectedCount)
}
