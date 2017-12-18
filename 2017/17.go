package main

import (
	"fmt"
)

func main() {
	currentPosition := 0
	afterZero := -1
	for i := 1; i <= 50000000; i++ {
		currentPosition = (currentPosition + 345)%i + 1
		if currentPosition == 1 {
			afterZero = i
		}
	}
	fmt.Println(afterZero)
}
