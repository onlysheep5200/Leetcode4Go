package main

import (
	"fmt"
)

/**
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

*/

type frame struct{
	nums []int
	left int
	start int
}


func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int,0, 100)
	if target < 0 || len(candidates) == 0{
		return result
	}

	stack := make([]frame, 0, 100)

	pop := func() frame {
		top := stack[len(stack) - 1]
		stack = stack[0:len(stack)-1]
		return top
	}

	newNums := func(nums []int, num int) []int{
		newNums := make([]int, len(nums), len(nums) + 1)
		copy(newNums, nums)
		newNums = append(newNums, num)
		return newNums
	}

	push := func(f frame) {

		for i := f.start; i<len(candidates); i++{
			val := candidates[i]
			left := f.left - val
			if left < 0 {
				continue
			}

			if left  == 0{
				result = append(result, newNums(f.nums, val))
				continue
			}

			stack = append(stack, frame{newNums(f.nums, val), left, i})
		}

	}

	push(frame{[]int{}, target, 0})

	for len(stack) != 0{
		push(pop())
	}

	return result
}

func main ()  {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}