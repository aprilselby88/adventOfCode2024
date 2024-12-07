package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseDecOne() ([]int, []int) {
	file, err := os.Open("input/decOne.txt")
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