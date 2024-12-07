package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_increasing_safely1(input_nums []int) bool {
	for i := 1; i < len(input_nums); i++ {
		increase := input_nums[i] - input_nums[i-1]
		if  increase > 3 || increase < 1 {
			return false
		}
	}
	return true
}

func is_decreasing_safely1(input_nums []int) bool {
	for i := 1; i < len(input_nums); i++ {
		decrease := input_nums[i] - input_nums[i-1]
		if decrease < -3 || decrease > -1 {
			return false
		}
	}
	return true
}


func is_decreasing_safely2(input_nums []int) bool {
	iter_nums := make([]int, len(input_nums))
	copy(iter_nums, input_nums)

	already_removed_bad_num := false
	stop_index := len(iter_nums)
	i := 1

	for i < stop_index {
		decrease := iter_nums [i] - iter_nums[i-1]
		current_num_is_bad := decrease < -3 || decrease > -1 

		if current_num_is_bad {
			// if already_removed_bad_num {
			// 	return false
			// }
			if (i == 1) {
				// handle first element removal
				new_iter_nums := iter_nums[1:]
				first_element_removed_is_safe := is_decreasing_safely1(new_iter_nums)
				if first_element_removed_is_safe {
					return true
				}
			} 

			// this doesn't handle the case where we have to remove the prev element to make current element valid

			// remove the bad number, adjust the stop index
			iter_nums = append(iter_nums[:i], iter_nums[i+1:]...)
			return is_decreasing_safely1(iter_nums)
			// already_removed_bad_num = true
			// stop_index--
		} else {
			i++
		}
	}
	if already_removed_bad_num {
		fmt.Println(input_nums)
		fmt.Println(iter_nums)
	}
	return true
}

func is_increasing_safely2(input_nums []int) bool {
	iter_nums := make([]int, len(input_nums))
	copy(iter_nums, input_nums)

	already_removed_bad_num := false
	stop_index := len(iter_nums)
	i := 1

	// doesn't handle first element removals

	for i < stop_index {
			increase := iter_nums[i] - iter_nums[i-1]
			current_num_is_bad := increase < 1 || increase > 3

			if current_num_is_bad {
					// if already_removed_bad_num {
					// 		return false
					// }

					// remove the bad number, adjust the stop index
					if (i == 1) {
						// handle first element removal
						new_iter_nums := iter_nums[1:]
						first_element_removed_is_safe := is_increasing_safely1(new_iter_nums)
						if first_element_removed_is_safe {
							return true
						}
					} 

					iter_nums = append(iter_nums[:i], iter_nums[i+1:]...)
					return is_increasing_safely1(iter_nums)
				
					// already_removed_bad_num = true
					// stop_index--
			} else {
					i++
			}
	}
	if already_removed_bad_num {
			fmt.Println(input_nums)
			fmt.Println(iter_nums)
	}
	return true
}

func is_decreasing_safely3(input_nums []int) bool {
	// Try removing each number one at a time
	for skip_index := 0; skip_index < len(input_nums); skip_index++ {
			// Create a new slice without the current number
			iter_nums := make([]int, 0, len(input_nums)-1)
			for i := 0; i < len(input_nums); i++ {
					if i != skip_index {
							iter_nums = append(iter_nums, input_nums[i])
					}
			}
			
			// Check if this removal creates a valid sequence
			if is_decreasing_safely1(iter_nums) {
					return true
			}
	}
	return false
}

func is_increasing_safely3(input_nums []int) bool {
	// Try removing each number one at a time
	for skip_index := 0; skip_index < len(input_nums); skip_index++ {
			// Create a new slice without the current number
			iter_nums := make([]int, 0, len(input_nums)-1)
			for i := 0; i < len(input_nums); i++ {
					if i != skip_index {
							iter_nums = append(iter_nums, input_nums[i])
					}
			}
			
			// Check if this removal creates a valid sequence
			if is_increasing_safely1(iter_nums) {
					return true
			}
	}
	return false
}


func main() {
	file, err := os.Open("input2.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	safeCount1 := 0
	safeCount2 := 0
	safeCount3 := 0
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Fields(line)
		nums := make([]int, len(numStrs))

		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}
			nums[i] = num
		}
		is_safe1 := is_increasing_safely1(nums) || is_decreasing_safely1(nums)
		is_safe2 := is_increasing_safely2(nums) || is_decreasing_safely2(nums)
		is_safe3 := is_increasing_safely3(nums) || is_decreasing_safely3(nums)

		if is_safe1 {
			safeCount1++
		}
		if is_safe2 {
			safeCount2++
		}
		if is_safe3 {
			safeCount3++
		}
		// if is_safe1 != is_safe2 {
		// 	fmt.Println("newly safe to part 2:", nums, "is_safe1:", is_safe1, "is_safe2:", is_safe2)
		// }
		if is_safe2 != is_safe3 {
			fmt.Println("newly safe to part 3:", nums, "is_safe2:", is_safe2, "is_safe3:", is_safe3)
		}
	}

	fmt.Println("safeCount1:", safeCount1)
	fmt.Println("safeCount2:", safeCount2)
	fmt.Println("safeCount3:", safeCount3)
}
