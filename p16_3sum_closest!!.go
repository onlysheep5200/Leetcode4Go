package main

import (
	"sort"
	"math"
)

/**
用双指针来省掉一次遍历
 */
func threeSumClosest(nums []int, target int) int {
	result := nums[0] + nums[1] + nums[len(nums) - 1]
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		left := i+1;
		right := len(nums) -1

		for ; left < right; {
			curSum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(curSum - target)) <  math.Abs(float64(result - target)){
				result = curSum
			}

			if curSum > target {
				right--
			}else if curSum < target{
				left ++
			}else {
				return result
			}
		}
	}

	return result
}