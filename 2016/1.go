package main

import (
    "fmt"
    "math"
    "strings"
    "strconv"
)

func main() {
    input := "R5, L2, L1, R1, R3, R3, L3, R3, R4, L2, R4, L4, R4, R3, L2, L1, L1, R2, R4, R4, L4, R3, L2, R1, L4, R1, R3, L5, L4, L5, R3, L3, L1, L1, R4, R2, R2, L1, L4, R191, R5, L2, R46, R3, L1, R74, L2, R2, R187, R3, R4, R1, L4, L4, L2, R4, L5, R4, R3, L2, L1, R3, R3, R3, R1, R1, L4, R4, R1, R5, R2, R1, R3, L4, L2, L2, R1, L3, R1, R3, L5, L3, R5, R3, R4, L1, R3, R2, R1, R2, L4, L1, L1, R3, L3, R4, L2, L4, L5, L5, L4, R2, R5, L4, R4, L2, R3, L4, L3, L5, R5, L4, L2, R3, R5, R5, L1, L4, R3, L1, R2, L5, L1, R4, L1, R5, R1, L4, L4, L4, R4, R3, L5, R1, L3, R4, R3, L2, L1, R1, R2, R2, R2, L1, L1, L2, L5, L3, L1"

    split := strings.Split(input, ", ")

    x := 0
    y := 0
    direction := "forwards"

    for i := 0; i < len(split); i++ {
        d := string(split[i][0])
        in :=  split[i][1:]
        num, _ := strconv.Atoi(in)

        dir, dX, dY := move(direction, d, num)
        direction = dir
        x = x + dX
        y = y + dY
    }

    fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func move(direction string, way string, length int) (d string, x int, y int) {
    switch (direction) {
    case "forwards":
        if way == "R" {
            d, x, y = "right", length, 0
        } else {
            d, x, y = "left", -length, 0
        }
        break
    case "backwards":
        if way == "R" {
            d, x, y = "left", -length, 0
        } else {
            d, x, y = "right", length, 0
        }
        break
    case "right":
        if way == "R" {
            d, x, y = "backwards", 0, -length
        } else {
            d, x, y = "forwards", 0, length
        }
        break
    case "left":
        if way == "R" {
            d, x, y = "forwards", 0, length
        } else {
            d, x, y = "backwards", 0, -length
        }
    }
    return
}
