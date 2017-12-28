package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	input = `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.`

	myInput = `Begin in state A.
Perform a diagnostic checksum after 12317297 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state D.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state C.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the right.
    - Continue with state F.

In state C:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state C.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.

In state D:
  If the current value is 0:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state E.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.

In state E:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the right.
    - Continue with state B.

In state F:
  If the current value is 0:
    - Write the value 0.
    - Move one slot to the right.
    - Continue with state C.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the right.
    - Continue with state E.`
)

type valueAction struct {
	write int
	direction int
	nextState string
}

type state struct {
	label string
	zero valueAction
	one valueAction
}

func main() {
	paragraphs := strings.Split(myInput, "\n\n")
	numOfRuns := getNumberOfRuns(paragraphs[0])
	pos := 0
	tape := []int{0}
	currentState := "A"
	states := map[string]state{}

	for _, p := range paragraphs[1:] {
		s := buildState(p)
		states[s.label] = s
	}

	for i := 0; i < numOfRuns; i++ {
		s := states[currentState]
		action := valueAction{}
		if tape[pos] == 0 {
			action = s.zero
		} else {
			action = s.one
		}
		tape[pos] = action.write
		pos += action.direction

		if pos > len(tape) - 1 {
			tape = append(tape, 0)
		} else if pos < 0 {
			tape = append([]int{0}, tape...)
			pos += 1
		}
		currentState = action.nextState
	}

	count := 0
	for _, n := range tape {
		if n == 1 {
			count++
		}
	}
	fmt.Println(count)
}

func getNumberOfRuns(p string) int {
	words := strings.Fields(strings.Split(p, "\n")[1])
	x, _ := strconv.Atoi(words[5])
	return x
}

func getLastWord(p string) string {
	fields := strings.Fields(p)
	lastWord := fields[len(fields)-1]
	return lastWord[:len(lastWord)-1]
}

func buildState(p string) state {
	lines := strings.Split(p, "\n")
	label := getLastWord(lines[0])
	return state{
		label,
		buildValueAction(lines[2:5]),
		buildValueAction(lines[6:9]),
	}
}

func buildValueAction(p []string) valueAction {
	write, _ := strconv.Atoi(getLastWord(p[0]))

	direction := 0
	if getLastWord(p[1]) == "right" {
		direction = 1
	} else {
		direction = -1
	}

	nextState := getLastWord(p[2])

	return valueAction{write, direction, nextState}
}
