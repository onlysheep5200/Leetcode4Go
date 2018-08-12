package main

/**
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
 */

 /**
  */
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	reachMap := make(map[int]bool)
	return canJumpFrom(&nums, 0, &reachMap)
}
func canJumpFrom(numsPtr *[]int, current int, reachMapPtr *map[int]bool) bool {
	if current == len(*numsPtr)-1 {
		return true
	}

	furthest := current + (*numsPtr)[current]
	if furthest >= len(*numsPtr) {
		return true
	}

	for i := 1; i <= (*numsPtr)[current]; i++ {
		reachable, exists := (*reachMapPtr)[current+i]
		if exists && reachable{
			return true
		}else if (canJumpFrom(numsPtr, current+i, reachMapPtr)){
			return true
		}
 	}

	(*reachMapPtr)[current] = false
	return false
}

func canJumpDP(nums []int) bool  {
	if len(nums) == 0 {
		return true
	}

	const REACHABLE int = 1
	const UNREACHABLE int = -1
	const UNKNOWN int = 0

	reachable := make([]int, len(nums))
	reachable[len(nums) - 1] = REACHABLE

	for i := len(nums) - 2; i >= 0; i-- {
		furthest := i+nums[i]
		if furthest > len(nums)-1{
			furthest = len(nums) - 1
		}

		for j := i + 1; j <= furthest; j++ {
			if reachable[j] == REACHABLE {
				reachable[i] = REACHABLE
				break
			}
		}
	}

	return reachable[0] == REACHABLE
}

func jumpGreedy(nums []int) bool {
	if len(nums)  == 0{
		return true
	}

	lastReachablePos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		furthest := i + nums[i]
		if furthest >= len(nums) {
			furthest = len(nums) - 1
		}

		if furthest  >= lastReachablePos {
			lastReachablePos = i
		}
	}

	return lastReachablePos == 0
}

func main() {
	//println(canJump([]int{2,3,1,1,4}))
	println(canJumpDP([]int{2,3,1,1,4}))
	println(canJumpDP([]int{3,2,1,0,4}))
}
