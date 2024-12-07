package main

import (
	"adventOfCode2024/input"
	solutions "adventOfCode2024/solutions"
	"fmt"
)

func main() {
    listOne, listTwo := input.ParseDecOne()
    fmt.Printf("input list one: %d", listOne)
    fmt.Printf("input list one: %d", listTwo)
    answer := solutions.DecOne(listOne, listTwo)
    fmt.Printf("solution: %d\n", answer)
}