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

func printGrid(grid [][]bool) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				fmt.Printf("%s", "#")
			} else {
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

func main() {
	grid := [][]bool{}
	for i, l := range strings.Split(input, "\n") {
		grid = append(grid, []bool{})
		for _, c := range strings.Split(l, "") {
			grid[i] = append(grid[i], c == "#")
		}
	}

	infectedCount := 0
	width := (len(grid)-1)/2
	t := turtle{width, width, "up"}
	for i := 0; i < 10000; i++ {
		if grid[t.x][t.y] {
			t.d = newDirection(t.d, "right")
			x, y := move(t)
			grid[t.x][t.y] = false
			t.x = t.x + x
			t.y = t.y + y
		} else {
			t.d = newDirection(t.d, "left")
			x, y := move(t)
			grid[t.x][t.y] = true
			infectedCount++
			t.x = t.x + x
			t.y = t.y + y
		}

		if t.x < 0 {
			grid = append([][]bool{[]bool{}}, grid...)
			for _ = range grid[1] {
				grid[0] = append(grid[0], false)
			}
			t.x = t.x + 1
		} else if t.x >= len(grid) {
			grid = append(grid, []bool{})
			for _ = range grid[1] {
				grid[len(grid)-1] = append(grid[len(grid)-1], false)
			}
		}

		if t.y < 0 {
			for i := range grid {
				grid[i] = append([]bool{false}, grid[i]...)
			}
			t.y = t.y + 1
		} else if t.y >= len(grid[0]) {
			for i := range grid {
				grid[i] = append(grid[i], false)
			}
		}
	}
//	printGrid(grid)
	fmt.Println(infectedCount)
}
