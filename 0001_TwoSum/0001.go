package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 6, 7, 11, 15}
	target := 9

	ss := twoSum(nums, target)
	fmt.Println(ss)
}

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
