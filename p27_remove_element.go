package main



func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	removedCount := 0
	length := len(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			removedCount++
			length--
		}else{
			nums[i - removedCount] = nums[i]
		}
	}

	return length
}