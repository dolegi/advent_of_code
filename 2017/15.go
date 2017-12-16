package main

import (
	"fmt"
)

type generator struct {
	factor int
	currentValue int
}

func (g *generator) nextValue() int {
	g.currentValue = (g.currentValue * g.factor) % 2147483647
	return g.currentValue
}

func main() {
	generatorA := &generator{16807, 873}
	generatorB := &generator{48271, 583}
	count := 0
	for i := 0; i < 40000000; i++ {
		a := fmt.Sprintf("%032b\n", generatorA.nextValue())
		b := fmt.Sprintf("%032b\n", generatorB.nextValue())

		if a[16:] == b[16:] {
			count++
		}
	}
	fmt.Printf("Win %d\n", count)
}
