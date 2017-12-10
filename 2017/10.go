package main

import (
	"fmt"
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

func main() {
	list := buildList(256)
	input := "130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224"
	lengths := append(getLengthsFromAscii(input), 17, 31, 73, 47, 23)
	sparseHash := calculateSparseHash(list, lengths)
	denseHash := calculateDenseHash(sparseHash)
	knotHash := hexEncodeArray(denseHash)
	fmt.Println(knotHash)
}
