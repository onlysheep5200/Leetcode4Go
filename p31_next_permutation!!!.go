package main

import (
	"fmt"
	"sort"
)

/**
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/


//找到第一个比前面一个数大的数，序号为i
//从nums[i]之后找到一个比nums[i]大且与其相差最小的数，与nums[i]交换
//令nums[i]之后的数增序排列（nums[i]之后本省是降序的）
func nextPermutation(nums []int)  {
	if nums == nil ||  len(nums)<=1 {
		return
	}

	subNums := nums

	for last := len(nums) - 1; last > 0; last -- {
		prevLast := last - 1
		if nums[prevLast] >= nums[last] {
			continue
		}

		minGap := nums[last] - nums[prevLast]
		replaceIdex := last

		for i := last+1; i < len(nums); i++ {
			if nums[i] <= nums[prevLast] {
				break
			}
			gap := nums[i] - nums[prevLast]
			if gap < minGap {
				minGap = gap
				replaceIdex = i
			}
		}

		nums[prevLast], nums[replaceIdex] = nums[replaceIdex], nums[prevLast]

		subNums = nums[last :]
		break
	}

	sort.Ints(subNums)
	//subLen := len(subNums)
	//for i := 0; i < len(subNums)/2; i++ {
	//	subNums[i], subNums[subLen - i - 1] = subNums[subLen - i - 1], subNums[i]
	//}
}

func main() {
	s := []int{1,3,2}
	nextPermutation(s)
	fmt.Println(s)
}