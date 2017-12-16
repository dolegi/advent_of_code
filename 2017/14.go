package main

import (
	"fmt"
	"strconv"
)

func buildList(size int) (list []int) {
	for i := 0; i < size; i++ {
		list = append(list, i)
	}
	return
}

func getItem(list []int, i int) int {
	return list[i%len(list)]
}

func getLengthsFromAscii(input string) (output []int) {
	for _, c := range input {
		output = append(output, int(c))
	}
	return
}

func calculateSparseHash(list []int, lengths []int) []int {
	currentPosition := 0
	skipSize := 0
	sparseHash := make([]int, len(list))
	for j := 0; j < 64; j++ {
		for _, length := range lengths {
			copy(sparseHash, list)
			for i := currentPosition; i < currentPosition+length; i++ {
				sparseHash[(length+currentPosition-i+currentPosition-1)%len(sparseHash)] = getItem(list, i)
			}
			currentPosition += length + skipSize
			skipSize++
			copy(list, sparseHash)
		}
	}
	return sparseHash
}

func calculateDenseHash(sparseHash []int) []int {
	xor := 0
	denseHash := make([]int, 0, 16)
	for i := 1; i <= len(sparseHash); i++ {
		xor ^= sparseHash[i-1]
		if i%16 == 0 {
			denseHash = append(denseHash, xor)
			xor = 0
		}
	}
	return denseHash
}

func hexEncodeArray(arr []int) (output string) {
	for _, n := range arr {
		output += fmt.Sprintf("%02x", n)
	}
	return
}

func knotHash(input string) string {
	lengths := append(getLengthsFromAscii(input), 17, 31, 73, 47, 23)
	sparseHash := calculateSparseHash(buildList(256), lengths)
	denseHash := calculateDenseHash(sparseHash)
	return hexEncodeArray(denseHash)
}

func appendNumber(input string, index int) string {
	num := strconv.Itoa(index)

	return input + "-" + num
}

func hexToBinary(input string) []bool {
	if len(input) > 8 {
		return append(hexToBinary(input[:8]), hexToBinary(input[8:])...)
	}
	n, err := strconv.ParseUint(input, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	binStr := fmt.Sprintf("%032b", n)

	bin := make([]bool, len(binStr))
	for i, b := range binStr {
		if b == '1' {
			bin[i] = true
		} else {
			bin[i] = false
		}
	}

	return bin
}

func countRegions(rows [][]bool) int {
	count := 0
	for i, row := range rows {
		for j, v := range row {
			if v {
				clearRegion(i, j, rows)
				count++
			}
		}
	}
	return count
}

func clearRegion(i, j int, rows [][]bool) {
	if i == -1 || i == len(rows) || j == -1 || j == len(rows[i]) || !rows[i][j] {
		return
	}
	rows[i][j] = false
	clearRegion(i-1, j, rows)
	clearRegion(i+1, j, rows)
	clearRegion(i, j+1, rows)
	clearRegion(i, j-1, rows)
}

func main() {
	input := "ljoxqyyw"

	rows := [][]bool{}
	for i := 0; i < 128; i++ {
		numberedInput := appendNumber(input, i)
		knotHash := knotHash(numberedInput)
		bin := hexToBinary(knotHash)
		rows = append(rows, bin)
	}
	fmt.Println(countRegions(rows))
}
