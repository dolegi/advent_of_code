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

func main() {
	list := buildList(256)
	fmt.Println(list)
	lengths := []int{130,126,1,11,140,2,255,207,18,254,246,164,29,104,0,224}
	currentPosition := 0
	skipSize := 0

	for _, length := range lengths {
		newList := make([]int, len(list))
		copy(newList, list)
		for i := currentPosition; i < currentPosition + length; i++ {
			newList[(length + currentPosition - i + currentPosition - 1)%len(newList)] = getItem(list, i)
		}
		currentPosition += length + skipSize
		skipSize++
		copy(list, newList)

		fmt.Println(newList)
	}
}
