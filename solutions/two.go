package adventofcode2024

func Two(reports [][]int) int {
	// 5 numbers in a line
	// safe : decrease by 1 or 2
	// safe: increase by 1, 2, 3
	// unsafe: increase then decrease of decrease then increase

	/*var r [][]int = [][]int{
        {7,6,4,2,1},
        {1,2,7,8,9},
		{9,7,6,2,1},
		{1,3,2,4,5},
		{8,6,4,4,1},
		{1,3,6,7,9},
    }*/

	unsafeReports := 0
	for _, report := range reports {
		increase := false
		decrease := false
		for i := 0; i < len(report)-1; i++ {
			// all increasing or all decreasing
			// two adjacent levels differ by at least one and at most three
			if(report[i] < report[i+1]){
				increase = true
				if (increase && decrease) {
					unsafeReports = unsafeReports + 1
					break
				}
				diff := report[i+1] - report[i]
				if (diff > 3) {
					unsafeReports = unsafeReports + 1
					break
				}
			} else if (report[i] == report[i+1]) {
				unsafeReports = unsafeReports + 1
				break
			} else if(report[i] > report[i+1]) {
				decrease = true
				if (increase && decrease) {
					unsafeReports = unsafeReports + 1
					break
				}
				diff := report[i] - report[i+1]
				if (diff > 3) {
					unsafeReports = unsafeReports + 1
					break
				}
			}
		}
	}

	return len(reports) - unsafeReports
}

func TwoB(reports [][]int) int {
	//if removing a single level from an unsafe report 
	//would make it safe, the report instead counts as safe

	/*var reports [][]int = [][]int{
        {7,6,4,2,1},
        {1,2,7,8,9},
		{9,7,6,2,1},
		{1,3,2,4,5},
		{8,6,4,4,1},
		{1,3,6,7,9},
    }*/

	unsafeReports := 0
	for _, report := range reports {
		increase := false
		decrease := false
		problemCount := 0
		for i := 0; i < len(report)-1; i++ {
			// all increasing or all decreasing
			// two adjacent levels differ by at least one and at most three
			if(report[i] < report[i+1]){
				increase = true
				if (increase && decrease) {
					problemCount = problemCount + 1
				}
				diff := report[i+1] - report[i]
				if (diff > 3) {
					problemCount = problemCount + 1
				}
			} else if (report[i] == report[i+1]) {
				problemCount = problemCount + 1
			} else if(report[i] > report[i+1]) {
				decrease = true
				if (increase && decrease) {
					problemCount = problemCount + 1
				}
				diff := report[i] - report[i+1]
				if (diff > 3) {
					problemCount = problemCount + 1
				}
			}
		}
		if (problemCount >= 1) {
			// remove one at a time and check
			removingFixes := false
			for j := 0; j < len(report); j++ {
				pC := 0
				increase := false
				decrease := false
				r := remove(report, j)
				for k := 0; k < len(r)-1; k++ {
					// all increasing or all decreasing
					// two adjacent levels differ by at least one and at most three
					if(r[k] < r[k+1]){
						increase = true
						if (increase && decrease) {
							pC = pC + 1
						}
						diff := r[k+1] - r[k]
						if (diff > 3) {
							pC = pC + 1
						}
					} else if (r[k] == r[k+1]) {
						pC = pC + 1
					} else if(r[k] > r[k+1]) {
						decrease = true
						if (increase && decrease) {
							pC = pC + 1
						}
						diff := r[k] - r[k+1]
						if (diff > 3) {
							pC = pC + 1
						}
					}
				}
				if (pC == 0) {
					removingFixes = true
					break
				}
			}
			if (!removingFixes) {
				unsafeReports = unsafeReports + 1
			}
		}
	}

	return len(reports) - unsafeReports
}

func remove(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
	x := append(ret, s[index+1:]...)
    return x 
}