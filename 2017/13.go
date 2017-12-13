package main

import (
	"fmt"
	"strconv"
	"strings"
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
	size            int
	scannerPosition int
	increase        bool
}

func moveScanners(delay int, limit int, lengths map[int]length) int {
	i := 0
	for i = 0; i < delay; i++ {
		for j := 0; j <= limit+1; j++ {
			currentLength := lengths[j]
			if currentLength.size == 0 {
				continue
			}

			increase := currentLength.increase
			if currentLength.scannerPosition > currentLength.size-2 {
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
		}
	}
	return i
}

func main() {
	lines := strings.Split(input, "\n")
	end, _ := strconv.Atoi(strings.Split(lines[len(lines)-1], ":")[0])
	limit := end + 1
	delay := 0
	originalLengths := map[int]length{}
	for _, line := range lines {
		lineArr := strings.Split(line, ":")
		id, _ := strconv.Atoi(lineArr[0])
		size, _ := strconv.Atoi(lineArr[1][1:])
		originalLengths[id] = length{size, 0, true}
	}

	for {
		lengths := map[int]length{}
		for key, value := range originalLengths {
			lengths[key] = value
		}

		i := 0
		for i = 0; i < limit; i++ {
			if lengths[i].size != 0 && 0 == lengths[i].scannerPosition {
				break
			}
			moveScanners(1, limit, lengths)
		}
		if i == limit {
			fmt.Println(delay)
			break
		}

		delay += 1
		moveScanners(1, limit, originalLengths)
	}
}
