package main

import (
	"fmt"
	"math"
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

var count int

func SieveOfEratosthenes(lower int, value int) {
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if !f[i] {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    for i := lower; i < value; i += 17 {
        if f[i] {
		count++
	}
    }
    fmt.Println(count)
}

func main() {
 	b := (79*100) + 100000
 	c := b + 17000
// 	d := 2
// 	e := 2
// 	f := 1
// 	h := 0

	SieveOfEratosthenes(b, c + 1)
// Loop:
// 	if (d * e) == b {
// 		f = 0
// 	}

// 	e += 1

// 	if e != b {
// 		goto Loop
// 	}

// 	d += 1

// 	if d != b {
// 		e = 2
// 		goto Loop
// 	}

// 	if f == 0 {
// 		h += 1
// 	}

// 	if b != c {
// 		b += 17
// 		f = 1
// 		d = 2
// 		e = 2
// 		goto Loop
// 	}

// 	fmt.Println(h)
}
