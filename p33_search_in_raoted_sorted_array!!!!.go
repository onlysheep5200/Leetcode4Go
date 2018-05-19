package main

import (
	"fmt"
)

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).
 */

func search(nums []int, target int) int {

	const INF int = int(^uint(0) >> 1)

	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	var mid int

	findCur := func() int {
		//target 和 nums[mid]在同一边，则按正常的二分来查找
		//否则，始终都将选取另一边
		//若target在往前翻的部分，mid在往后翻的部分，则将分区左移，故返回INF
		//若target在往后翻的部分，mid在往前翻的部分，则将分区右移，故返回-INF
		if (target >= nums[0]) == (nums[mid] >= nums[0]) {
			return nums[mid]
		}

		if target >= nums[0] {
			return INF
		}else {
			return -INF
		}
	}


	for left < right {
		mid = (left + right) / 2

		cur := findCur()

		if cur == target {
			return mid
		}else if cur > target{
			right = mid - 1
		}else {
			left = mid + 1
		}
	}

	if left >= 0 && left < len(nums) && nums[left] == target {
		 return left
	}

	if right >= 0 && right < len(nums) && nums[right] == target {
		return right
	}

	return -1
}

func main() {
	fmt.Println(search([]int{1,3}, 4))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}