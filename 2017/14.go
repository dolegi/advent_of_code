package main

import (
	"fmt"
	"strconv"
	"strings"
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
		if i % 16 == 0 {
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

func hexToBinary(input string) string {
	if len(input) > 8 {
		return hexToBinary(input[:8]) + hexToBinary(input[8:])
	}
	n, err := strconv.ParseUint(input, 16, 32)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%b", n)
}

func countSquares(input string) int {
	count := 0
	for _, c := range strings.Split(input, "") {
		if c == "1" {
			count++
		}
	}
	return count
}

func main() {
	input := "ljoxqyyw"

	count := 0
	for i := 0; i < 128; i++ {
		numberedInput := appendNumber(input, i)
		knotHash := knotHash(numberedInput)
		bin := hexToBinary(knotHash)
		count += countSquares(bin)
	}

	fmt.Println(count)
}
