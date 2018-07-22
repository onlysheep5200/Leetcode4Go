package main

import "fmt"

/**
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:

Input: [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
    Jump 1 step from index 0 to 1, then 3 steps to the last index.
Note:

You can assume that you can always reach the last index.

 */

 /**
 贪心
 主要概念：
 	最大覆盖范围
 	当前终点
 	跳数
 当当前所在位置 > 当前终点，且当前所在位置没有达到最后一个点时，需要跳一步
  */
func jump(nums []int) int {
	curEnd := 0
	curMaxRange := 0
	jumps := 0
	for i, v := range nums{
		if i+v > curMaxRange {
			curMaxRange = i+v
		}

		if i>= curEnd && i < len(nums) - 1{
			curEnd = curMaxRange
			jumps++
		}
	}
	return jumps
}



func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
}