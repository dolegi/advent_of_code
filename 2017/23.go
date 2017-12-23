package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	input = `set b 79
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23`
)

var registers map[string]int
var multiCount int

func getValue(a string) int {
	if a == "" {
		return 0
	} else if num, err := strconv.Atoi(a); err == nil {
		return num
	}
	return registers[a]
}

func runCmd(cmd, a string, b int) int {
	switch cmd {
	case "set":
		registers[a] = b
	case "sub":
		registers[a] -= b
	case "mul":
		registers[a] *= b
		multiCount = multiCount + 1
	case "jnz":
		if num, err := strconv.Atoi(a); err == nil && num != 0 {
			return b
		} else if registers[a] != 0 {
			return b
		}
	}
	return 1
}

func destructureLine(line string) (cmd, a, b string) {
	data := strings.Split(line, " ")
	if len(data) == 2 {
		return data[0], data[1], ""
	}
	return data[0], data[1], data[2]
}

func main() {
	registers = make(map[string]int)
	lines := strings.Split(input, "\n")
	offset := 0
	multiCount = 0
	for i := 0; i < len(lines); i += offset {
		cmd, a, b := destructureLine(lines[i])
		offset = runCmd(cmd, a, getValue(b))
//		fmt.Println(i, cmd, a, getValue(b), registers)
	}
	fmt.Println(multiCount)
}
