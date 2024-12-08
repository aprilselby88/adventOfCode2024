package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseOne() ([]int, []int) {
	file, err := os.Open("input/one.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	var listOne []int
	var listTwo []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")
		numOne, _ := strconv.Atoi(nums[0])
		numTwo, _ := strconv.Atoi(nums[1])
		listOne = append(listOne, numOne)
		listTwo = append(listTwo, numTwo)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}

	return listOne, listTwo
}

func ParseTwo() ([][]int) {
	file, err := os.Open("input/two.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		var levels []int
		for _, val := range nums {
			num, _ := strconv.Atoi(val)
			levels = append(levels, num)
		}
		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return reports
}