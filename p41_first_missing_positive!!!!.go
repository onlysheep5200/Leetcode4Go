package main

import "fmt"

/**
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:

Input: [1,2,0]
Output: 3
Example 2:

Input: [3,4,-1,1]
Output: 2
Example 3:

Input: [7,8,9,11,12]
Output: 1
Note:

Your algorithm should run in O(n) time and uses constant extra space.
 */

 /**
 一个很牛逼的swap方法：三次异或
 public void swap(int[] A, int i, int j){
    if(i!=j){
        A[i]^=A[j];
        A[j]^=A[i];
        A[i]^=A[j];
    }
}
  */
func firstMissingPositive(nums []int) int {
	//方法1：把nums[i]与nums[nums[i]-1]交换；第一个nums[i] != i 的i+1即是缺失的正数
	//Put each number in its right place.
	if len(nums) == 0{
		return 1
	}

	for i := 0; i < len(nums);  {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i] - 1] != nums[i] {
			nums[nums[i] - 1], nums[i] = nums[i],nums[nums[i]-1]
		}else{
			i++
		}
	}

	for i := 0; i<len(nums); i++  {
		if nums[i] != i+1 {
			return i+1
		}
	}

	//如果nums的各个数顺序都是对的，那么缺失的一定是len(nums) ,因为nums数组最大只到len(nums)-1
	return len(nums)+1
}

func main() {
	fmt.Println(firstMissingPositive([]int{6, 3, 4, 1, 8, 2}))
}