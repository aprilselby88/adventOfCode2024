package adventofcode2024

func OneB(listOne []int, listTwo []int) int {
	score := 0
	
	// for every number in left list, count how many times it is in the right list
	// build a hashmap of numbers and counts of the second list
	// loop through first list and look up in hashmap
	f := make(map[int]int)
	for i := 0; i < len(listTwo); i++ {
		key := listTwo[i]
		if val, found := f[key]; found {
			f[key] = val + 1
		} else {
			f[key] = 1
		}
	}

	for i := 0; i < len(listOne); i++ {
		key := listOne[i]
		if _, found := f[key]; found {
			score = score + (f[key] * key)
		}
	}

	return score
}