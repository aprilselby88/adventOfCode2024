package adventofcode2024

import (
	"regexp"
	"strconv"
)

func Three(input string) int {
	total := 0
	
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	reDig := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(input, -1)

	on := true
	for _, match := range matches {
		// nums
		digs := reDig.FindAllString(match, -1)
		if (len(digs) != 2) {
			// do or don't
			if (match == "do()") {
				on = true
			} else if (match == "don't()") {
				on = false
			}
		} else {
			if on {
				digitOne, _ := strconv.Atoi(string(digs[0]))
				digitTwo, _ := strconv.Atoi(string(digs[1]))
				total += (digitOne * digitTwo)
			}
		}
	}

	return total
}