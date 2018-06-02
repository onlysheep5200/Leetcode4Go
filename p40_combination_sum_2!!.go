package main

import "sort"

/**
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
*/

//same as p39

type frame struct{
	nums []int
	left int
	start int
}


func combinationSum2(candidates []int, target int) [][]int {

	result := make([][]int,0, 100)
	if target < 0 || len(candidates) == 0{
		return result
	}

	stack := make([]frame, 0, 100)
	sort.Ints(candidates)

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

			if i > f.start && val == candidates[i-1] {
				continue
			}


			left := f.left - val
			if left < 0 {
				continue
			}

			if left  == 0{
				result = append(result, newNums(f.nums, val))
				continue
			}

			stack = append(stack, frame{newNums(f.nums, val), left, i+1})
		}

	}

	push(frame{[]int{}, target, 0})

	for len(stack) != 0{
		push(pop())
	}

	return result
}



