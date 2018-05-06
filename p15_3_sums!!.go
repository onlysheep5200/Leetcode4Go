package main

import (
	"sort"
	"fmt"
)
//使用双指针发减少一次遍历
func threeSum(nums []int) [][]int {
	result := make([][]int, 0, 100)
	if nums == nil || len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	contains := make(map[string]bool)

	for i:=0; i<len(nums); i++  {

		if nums[i] >0 {
			break
		}

		left := i+1
		right := len(nums) - 1

		for ; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				key := fmt.Sprintf("%d_%d_%d", nums[i], nums[left], nums[right])
				if _, contain := contains[key]; !contain {
					contains[key] = true
					result = append(result, []int{nums[i], nums[left], nums[right]})
				}
				left ++
				right --
			}else if sum > 0{
				right --
			}else {
				left ++
			}

		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}