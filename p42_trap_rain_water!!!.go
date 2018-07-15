package main

import "fmt"

/**
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
 */

func trap(height []int) int {

	return 0
}

/**
利用栈
若当前高度大于栈顶高度，则说明当前高度与栈顶之前的条块形成了夹层，此时可以盛水
 */
func trapByStack(height []int) int{
	if len(height) == 0 {
		return 0
	}

	stack := make([]int,0,len(height))
	total := 0

	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[stack[len(stack)-1]] < height[i] {
			mid := stack[len(stack) - 1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				break
			}

			minHeight := height[i]
			top := stack[len(stack) - 1]
			if height[top] < height[i] {
				minHeight = height[top]
			}
			total += (minHeight - height[mid]) * (i - top - 1)
		}
		stack = append(stack, i)
	}
	return total
}

func main() {
	fmt.Println(trapByStack([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}