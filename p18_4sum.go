package main

import (
	"sort"
	"fmt"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 100)
	sort.Ints(nums)
	candidate := make(map[string]byte)

	for i := 0; i < len(nums); i++ {
		leftOf := nums[i]
		leftTarget := target - leftOf
		for j := i+1; j < len(nums); j++ {

			first := nums[j]
			left := j+1
			right := len(nums) - 1

			for ; left < right; {
				sum := first + nums[left] + nums[right]
				if sum == leftTarget {
					key := fmt.Sprintf("%d_%d_%d_%d", leftOf, first, nums[left], nums[right])
					if _, exists := candidate[key]; !exists {
						candidate[key] = 1
						result = append(result, []int{leftOf, first, nums[left], nums[right]})
					}
					left ++
					right --
				}else if sum > leftTarget{
					right --
				}else {
					left ++
				}
			}
		}
	}
	return result
}

