package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)



func main() {
	file, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var leftNumbers []int
	var rightNumbers []int
	rightNumberCounts := make(map[int]int) 

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		nums := strings.Fields(line)
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error converting left number:", err)
			return
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println("Error converting right number:", err)
			return
		}
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)

		rightNumberCounts[right]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	file.Close()

	// part 1

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	sumDeltas := 0
	for i, _ := range leftNumbers {
		delta := leftNumbers[i] - rightNumbers[i]
		if delta > 0 {
			sumDeltas += delta
		} else {
			sumDeltas -= delta
		}
	}

	fmt.Println("Part 1:",sumDeltas)

	countSum := 0

	// part 2
	for _, n := range leftNumbers {
		countSum += n * rightNumberCounts[n]
	}

	fmt.Println("Part 2:",countSum)
}