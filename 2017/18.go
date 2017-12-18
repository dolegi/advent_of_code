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

type program struct {
	registers map[string]int
	queue []int
	terminated bool
	location int
	locked bool
	sendCount int
}

func getValue(a string, p program) int {
	if a == "" {
		return 0
	} else if num, err := strconv.Atoi(a); err == nil {
		return num
	}
	return p.registers[a]
}

func runCmd(cmd, a string, b, limit int, one, two *program) {
	switch cmd {
	case "set":
		one.registers[a] = b
	case "add":
		one.registers[a] += b
	case "mul":
		one.registers[a] *= b
	case "mod":
		one.registers[a] %= b
	case "snd":
		one.sendCount += 1
		if num, err := strconv.Atoi(a); err == nil {
			two.queue = append(two.queue, num)
		} else {
			two.queue = append(two.queue, one.registers[a])
		}
	case "rcv":
		if len(one.queue) == 0 {
			one.locked = true
			return
		}
		one.locked = false
		one.registers[a] = one.queue[0]
		one.queue = one.queue[1:]
	case "jgz":
		if num, err := strconv.Atoi(a); err == nil && num > 0 {
			one.location += b
			return
		} else if one.registers[a] > 0 {
			one.location += b
			return
		}
	}
	one.location += 1
}

func destructureLine(line string) (cmd, a, b string) {
	data := strings.Split(line, " ")
	if len(data) == 2 {
		return data[0], data[1], ""
	}
	return data[0], data[1], data[2]
}

func main() {
	p1 := program{map[string]int{"p":0}, []int{}, false, 0, false, 0}
	p2 := program{map[string]int{"p":1}, []int{}, false, 0, false, 0}

	lines := strings.Split(input, "\n")
	for !((p1.terminated && p2.terminated) || (p1.locked && p2.locked)) {
		if p1.location >= len(lines) -1 || p1.location < 0 {
			p1.terminated = true
		} else {
			cmd, a, b := destructureLine(lines[p1.location])
			runCmd(cmd, a, getValue(b, p1), len(lines), &p1, &p2)
		}

		if p2.location >= len(lines)-1 || p2.location < 0{
			p2.terminated = true
		} else {
			cmd, a, b := destructureLine(lines[p2.location])
			runCmd(cmd, a, getValue(b, p2), len(lines), &p2, &p1)
		}
	}
	fmt.Println(p1.sendCount)
	fmt.Println(p2.sendCount)
}
