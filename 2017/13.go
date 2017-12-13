package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	input2 = `0: 3
1: 2
4: 4
6: 4`
	input = `0: 3
1: 2
2: 4
4: 8
6: 5
8: 6
10: 6
12: 4
14: 6
16: 6
18: 9
20: 8
22: 8
24: 8
26: 8
28: 10
30: 8
32: 12
34: 10
36: 14
38: 12
40: 12
42: 12
44: 12
46: 12
48: 12
50: 14
52: 12
54: 14
56: 12
58: 12
60: 14
62: 18
64: 14
68: 14
70: 14
72: 14
74: 14
78: 14
80: 20
82: 14
84: 14
90: 17`
)

type length struct {
	size int
	scannerPosition int
	increase bool
}

func main() {
	lines := strings.Split(input, "\n")
	end, _ := strconv.Atoi(strings.Split(lines[len(lines)-1], ":")[0])
	lengths := make(map[int]length)

	for _, line :=  range lines {
		lineArr := strings.Split(line, ":")
		id, _ := strconv.Atoi(lineArr[0])
		size, _ := strconv.Atoi(lineArr[1][1:])
		lengths[id] = length{size, 0, true}
	}

	severity := 0
	for i := 0; i <= end; i++ {
		if lengths[i].size != 0 && 0 == lengths[i].scannerPosition {
			severity += (i * lengths[i].size)
			fmt.Printf("HIT %d %d\n", i, lengths[i].size)
		}
//		fmt.Printf("%d %d | ", i, lengths[i].scannerPosition)
		for j := 0; j <= end; j++ {
			currentLength := lengths[j]
			if currentLength.size == 0 {
				continue
			}

			increase := currentLength.increase
			if currentLength.scannerPosition > currentLength.size - 2 {
				increase = false
			}
			if currentLength.scannerPosition == 0 {
				increase = true
			}
			scannerPosition := currentLength.scannerPosition
			if increase {
				scannerPosition += 1
			} else {
				scannerPosition -= 1
			}
			lengths[j] = length{currentLength.size, scannerPosition, increase}
//			fmt.Printf("%d %v   ", j, currentLength)
		}
//		fmt.Println()
	}
	fmt.Println(severity)
}
