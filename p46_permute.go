package main

import "fmt"

func permute(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0, 100)
	visited := make([]bool, len(nums))
	result = nextPermute(nums,make([]int, 0, len(nums)), visited, result)
	return result
}

func nextPermute(nums []int, candidate []int, visited []bool, result [][]int) [][]int {
	if len(candidate) == len(nums) {
		newPermutate := make([]int, len(nums))
		copy(newPermutate, candidate)
		result = append(result, newPermutate)
		return result
	}
	for i := 0; i < len(nums); i++ {
		if !(visited[i]) {
			visited[i] = true
			candidate = append(candidate, nums[i])
			result = nextPermute(nums, candidate,visited, result)
			visited[i] = false
			candidate = candidate[0:len(candidate)-1]
		}
	}
	return result
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
}