package main

import (
	"adventOfCode2024/input"
	solutions "adventOfCode2024/solutions"
	"fmt"
)

func main() {
	listOne, listTwo := input.ParseOne()
	answer := solutions.OneA(listOne, listTwo)
	fmt.Printf("solution one A: %d\n", answer)
    answer = solutions.OneB(listOne, listTwo)
    fmt.Printf("solution one B: %d\n", answer)

    listThree := input.ParseTwo()
    answer = solutions.Two(listThree)
    fmt.Printf("solution two A: %d\n", answer)
    answer = solutions.TwoB(listThree)
    fmt.Printf("solution two B: %d\n", answer)
}
