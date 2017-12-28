package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	input = `50/41
19/43
17/50
32/32
22/44
9/39
49/49
50/39
49/10
37/28
33/44
14/14
14/40
8/40
10/25
38/26
23/6
4/16
49/25
6/39
0/50
19/36
37/37
42/26
17/0
24/4
0/36
6/9
41/3
13/3
49/21
19/34
16/46
22/33
11/6
22/26
16/40
27/21
31/46
13/2
24/7
37/45
49/2
32/11
3/10
32/49
36/21
47/47
43/43
27/19
14/22
13/43
29/0
33/36
2/6`
)

var maxStrength int
var maxLength int

type branch struct {
	id int
	a int
	b int
	hit bool
}

func main() {
	branches := buildBranches(input)
	fmt.Println(branches)

	findMaxLength(0, 0, 0, branches)
	fmt.Println(maxStrength)
}

func buildBranches(in string) []branch {
	branches := []branch{}
	lines := strings.Split(in, "\n")
	for i, line := range lines {
		values := strings.Split(line, "/")
		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])
		branches = append(branches, branch{i, a, b, false})
	}
	return branches
}

func findMaxLength(node, strength, length int, branches []branch) {
	if length >= maxLength {
		maxLength = length
		if maxStrength < strength {
			maxStrength = strength
		}
	}

	for i, brch := range branches {
		if brch.hit {
			continue
		}
		if brch.a == node {
			branches[i].hit = true
			findMaxLength(brch.b, strength+brch.a+brch.b, length+1, branches)
			branches[i].hit = false
		}
		if brch.b == node {
			branches[i].hit = true
			findMaxLength(brch.a, strength+brch.a+brch.b, length+1, branches)
			branches[i].hit = false
		}
	}
}
