package main

import (
	"adventOfCode2024/input"
	solutions "adventOfCode2024/solutions"
	"fmt"
)

func main() {
	listOne, listTwo := input.Parse()
	answer := solutions.One(listOne, listTwo)
	fmt.Printf("solution one: %d\n", answer)

    answer = solutions.Two(listOne, listTwo)
    fmt.Printf("solution two: %d\n", answer)
}
