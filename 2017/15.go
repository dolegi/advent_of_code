package main

import (
	"fmt"
)

type generator struct {
	factor int
	currentValue int
	criteria int
}

func (g *generator) nextValue() int {
	g.currentValue = (g.currentValue * g.factor) % 2147483647
	if g.currentValue % g.criteria == 0 {
		return g.currentValue
	} else {
		return g.nextValue()
	}
}

func main() {
	generatorA := &generator{16807, 873, 4}
	generatorB := &generator{48271, 583, 8}
	count := 0
	for i := 0; i < 5000000; i++ {
		a := fmt.Sprintf("%032b\n", generatorA.nextValue())
		b := fmt.Sprintf("%032b\n", generatorB.nextValue())

		if a[16:] == b[16:] {
			count++
		}
	}
	fmt.Printf("Win %d\n", count)
}
