package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 1)

	currentPosition := 0
	for i := 1; i < 2018; i++ {
		currentPosition = (currentPosition + 345)%len(arr) + 1
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		arr = append(append(arr[:currentPosition], i), arrCopy[currentPosition:]...)
	}
	fmt.Println(arr[currentPosition + 1])
}
