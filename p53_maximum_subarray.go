package main

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 */

func maxSubArray(nums []int) int {
	sum := ^int(^uint32(0)>>1)
	prevSum := 0

	if len(nums) == 0 {
		return sum
	}

	for i := 0; i < len(nums); i++ {
		prevSum += nums[i]
		if prevSum > sum {
			sum = prevSum
		}
		if prevSum < 0 {
			prevSum = 0
		}
	}

	return sum
}

func main() {
	println(maxSubArray([]int{-2,-1,-3}))
}
