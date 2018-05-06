package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for ; left < right; {
		minHeight := height[left]
		if height[right] < height[left] {
			minHeight = height[right]
		}
		area := (right - left) * minHeight
		if area > maxArea {
			maxArea = area
		}

		//长度取决于最小的那条高，且移动指针会导致宽度变短以致面积减小。
		// 因此需移动较短的那条高，以寻求更大的面积来抵消宽度的减小并获取更大的高度
		if minHeight == height[left] {
			left++
		}else{
			right--
		}
	}
	return maxArea
}

func main() {
	var t []int = []int{1,2,3,4}
	fmt.Println(maxArea(t))
}
