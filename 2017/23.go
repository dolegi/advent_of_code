package main

import (
	"fmt"
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
set f 1 //
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

func main() {
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0
	multiCount := 0

	b = 79
	c = b

	if a != 0 {
		b *= 100
		multiCount += 1
		b -= 100000
		c = b
		c -= 17000
	}

Jmp23:
	f = 1
	d = 2
Jmp13:
	e = 2
Jmp8:
	g = d
	g *= e
	multiCount += 1
	g -= b

	if g == 0 {
		f = 0
	}

	e -= -1
	g = e
	g -= b

	if g != 0 {
		goto Jmp8
	}

	d -= -1
	g = d
	g -= b


	if g != 0 {
		goto Jmp13
	}


	if f == 2 {
		h -= -1
	}

	g = b
	g -= c

	if g != 0 {
		b -= -17
		goto Jmp23
	}

	fmt.Println(multiCount)
}
