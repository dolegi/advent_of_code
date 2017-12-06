package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	input = "4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3"
	Limit = int(^uint(0) >> 1)
)

func StrToIntArray(arr string) (resultArr []int) {
	strArr := strings.Fields(arr)
	resultArr = make([]int, len(strArr))
	for i, str := range strArr {
		resultArr[i], _ = strconv.Atoi(str)
	}
	return
}

func equalSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func sliceInMatrix(slice []int, matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		if equalSlices(slice, matrix[i]) {
			return true
		}
	}
	return false
}

func findHighest(arr []int) (index int) {
	for i := range arr {
		if arr[i] > arr[index] {
			index = i
		}
	}
	return
}

func calculateCyclesInLoop(currentConfig []int) int {
	prevConfigs := [][]int{}

	for i := 0; !sliceInMatrix(currentConfig, prevConfigs); i++ {
		if i >= Limit {
			return -1
		}
		config := make([]int, len(currentConfig))
		copy(config, currentConfig)
		prevConfigs = append(prevConfigs, config)

		index := findHighest(currentConfig)
		value := currentConfig[index]
		currentConfig[index] = 0
		for j := 0; j < value; j++ {
			location := (index + 1 + j) % len(currentConfig)
			currentConfig[location] += 1
		}
	}

	startOfLoop := 0
	for i := 0; i < len(prevConfigs); i++ {
		if equalSlices(currentConfig, prevConfigs[i]) {
			startOfLoop = i
			continue
		}
	}

	return len(prevConfigs) - startOfLoop
}

func main() {
	inputArr := StrToIntArray(input)
	fmt.Println(calculateCyclesInLoop(inputArr))
}
