package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	input = `set i 31
set a 1
mul p 17
jgz p p
mul a 2
add i -1
jgz i -2
add a -1
set i 127
set p 464
mul p 8505
mod p a
mul p 129749
add p 12345
mod p a
set b p
mod b 10000
snd b
add i -1
jgz i -9
jgz a 3
rcv b
jgz b -1
set f 0
set i 126
rcv a
rcv b
set p a
mul p -1
add p b
jgz p 4
snd a
set a b
jgz 1 3
snd b
set f 1
add i -1
jgz i -11
snd a
jgz f -16
jgz a -19`
)

var registers map[string]int
var sounds map[string]int

func getValue(a string) int {
	if a == "" {
		return 0
	} else if num, err := strconv.Atoi(a); err == nil {
		return num
	}
	return registers[a]
}

func runCmd(cmd, a string, b, limit int) int {
	switch cmd {
	case "set":
		registers[a] = b
	case "add":
		registers[a] += b
	case "mul":
		registers[a] *= b
	case "mod":
		registers[a] %= b
	case "snd":
		sounds[a] = registers[a]
	case "rcv":
		if registers[a] != 0 {
			registers[a] = sounds[a]
			fmt.Println(sounds)
			return limit
		}
	case "jgz":
		if registers[a] != 0 {
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
	sounds = make(map[string]int)
	lines := strings.Split(input, "\n")
	offset := 0
	for i := 0; i < len(lines); i += offset {
		cmd, a, b := destructureLine(lines[i])
		offset = runCmd(cmd, a, getValue(b), len(lines))
	}
}
