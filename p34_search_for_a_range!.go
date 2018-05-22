package main

//可以有更优化的方案
//优化：找到第一个target 和第一个 target+1 的index即可求出空间
func searchRange(nums []int, target int) []int {
	left := -1
	right := -1
	if len(nums) == 0 {
		return []int{left, right}
	}

	type wrapper struct {
		slice []int;
		start int;
	}

	space := make([]wrapper, 0, len(nums))
	space = append(space, wrapper{nums, 0})


	findTargetRange := func(w wrapper){
		l := 0
		r := len(w.slice) - 1

		for l <= r {
			mid := (l + r) / 2
			if w.slice[mid] == target {
				if left == -1 && right == -1 {
					left = mid + w.start
					right = mid + w.start
				}else{
					if left > mid + w.start {
						left = mid + w.start
					}else if right < mid + w.start {
						right = mid + w.start
					}
				}
				if mid != 0 {
					space = append(space, wrapper{w.slice[0 : mid], w.start})
				}
				if mid != len(w.slice) - 1 {
					space = append(space, wrapper{w.slice[mid+1 : ], w.start + mid + 1})
				}
				return
			}else if target >  w.slice[mid]{
				l = mid + 1
			}else {
				r = mid - 1
			}
		}
	}

	for len(space) != 0  {
		findTargetRange(space[0])
		space = space[1:]
	}

	return []int{left, right}
}
