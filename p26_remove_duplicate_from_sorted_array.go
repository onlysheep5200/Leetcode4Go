package main


func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	p := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev{
			nums[p] = nums[i]
			p++
			prev = nums[i]
		}
	}
	return p
}