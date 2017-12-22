package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	rules = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`
)

type rule struct {
	before [][]byte
	after  [][]byte
	width  int
}

type coord struct {
	x int
	y int
}

func transform(pattern [][]byte) {
	for i := range pattern {
		for j := i; j < len(pattern[i]); j++ {
			keep := pattern[i][j]
			pattern[i][j] = pattern[j][i]
			pattern[j][i] = keep
		}
	}
}

func reverseColumns(pattern [][]byte) {
	for i := range pattern {
		k := len(pattern) - 1
		for j := 0; j < k; j++ {
			keep := pattern[j][i]
			pattern[j][i] = pattern[k][i]
			pattern[k][i] = keep
			k--
		}
	}
}

func reverseArr(a []byte) []byte {
	if len(a) == 0 {
		return a
	}
	return append(reverseArr(a[1:]), a[0])
}

func rotate90(pattern [][]byte) {
	transform(pattern)
	reverseColumns(pattern)
}

func flipLeftRight(pattern [][]byte) {
	for i := range pattern {
		pattern[i] = reverseArr(pattern[i])
	}
}

func flipTopBottom(pattern [][]byte) {
	transform(pattern)
	flipLeftRight(pattern)
	transform(pattern)
}

func printMatrix(m [][]byte) {
	for _, l := range m {
		fmt.Println(string(l))
	}
	fmt.Println()
}

func flipFlop(pattern [][]byte) {
	for i := 0; i < 4; i++ {
		printMatrix(pattern)
		flipLeftRight(pattern)
		printMatrix(pattern)
		flipTopBottom(pattern)
		printMatrix(pattern)

		flipLeftRight(pattern)
		flipTopBottom(pattern)

		rotate90(pattern)
	}
}

func strToByteMatrix(str string) [][]byte {
	bytes := [][]byte{[]byte{}}
	level := 0
	for _, r := range str {
		if r == '/' {
			level++
			bytes = append(bytes, []byte{})
			continue
		}
		bytes[level] = append(bytes[level], byte(r))
	}
	return bytes
}

func strToRule(str string) rule {
	split := strings.Split(str, " ")
	before := strToByteMatrix(split[0])
	after := strToByteMatrix(split[2])
	return rule{before, after, len(before)}
}

func printByteMatrix(b [][]byte) {
	for i := range b {
		fmt.Printf("%s\n", string(b[i]))
	}
	fmt.Println()
}

func compareByteMatrices(a [][]byte, b [][]byte) bool {
	for i := range a {
		if bytes.Compare(a[i], b[i]) != 0 {
			return false
		}
	}
	return true
}

func compareWithRule(pattern [][]byte, r rule) ([][]byte, bool) {
	for i := 0; i < 4; i++ {
		if compareByteMatrices(pattern, r.before) {
			return r.after, true
		}
		flipLeftRight(pattern)
		if compareByteMatrices(pattern, r.before) {
			return r.after, true
		}
		flipTopBottom(pattern)
		if compareByteMatrices(pattern, r.before) {
			return r.after, true
		}
		rotate90(pattern)
	}

	return pattern, false
}

func transformPattern(pattern [][]byte, programRules []rule) [][]byte {
	for _, r := range programRules {
		if len(pattern)%r.width != 0 {
			continue
		}
		if len(pattern) > r.width {
			newWidth := len(pattern) + (len(pattern) / r.width)
			fmt.Println("WIDTH", newWidth)
		}
		newPattern, changed := compareWithRule(pattern, r)
		if changed {
			return newPattern
		}
	}
	return pattern
}

func splitPattern(pattern [][]byte, width int) [][][]byte {
	patternArr := [][][]byte{}

	coords := generateSplitCoords(len(pattern), width)
	for _, c := range coords {
		nextPattern := [][]byte{}
		for i := 0; i < width; i++ {
			nextPattern = append(nextPattern, []byte{})
			for j := 0; j < width; j++ {
				nextPattern[i] = append(nextPattern[i], pattern[c.x + i][c.y + j])
			}
		}
		patternArr = append(patternArr, nextPattern)
	}

	return patternArr
}

func generateSplitCoords(num int, max int) []coord {
	coords := []coord{}
	nums := []int{}
	factor := num / max
	for i := 1; i <= factor; i++ {
		nums = append(nums, num - (max*i))
	}

	for _, i := range nums {
		for _, j := range nums {
			coords = append([]coord{coord{i,j}}, coords...)
		}
	}
	return coords
}

func main() {
	pattern := [][]byte{
		[]byte{'.', '#', '.'},
		[]byte{'.', '.', '#'},
		[]byte{'#', '#', '#'},
	}

	programRules := []rule{}
	for _, r := range strings.Split(rules, "\n") {
		programRules = append(programRules, strToRule(r))
	}

	printByteMatrix(pattern)
	pattern = transformPattern(pattern, programRules)
	printByteMatrix(pattern)

	patterns := splitPattern(pattern, 2)
	for _, p := range patterns {
		printByteMatrix(p)
	}
}
