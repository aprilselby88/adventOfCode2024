package adventofcode2024

import (
	"sort"
)

func OneA(listOne []int, listTwo []int) int {
	sort.Ints(listOne)
	sort.Ints(listTwo)

	difference := 0
	for i := 0; i < len(listOne); i++ {
		diff := listOne[i] - listTwo[i]
		if diff < 0 {
			diff = -diff
		}
		difference = difference + diff
	}

	return difference
}
