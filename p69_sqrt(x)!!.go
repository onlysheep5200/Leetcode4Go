package main

/**
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
 */

func mySqrt(x int) int {
	if x == 1{
		return x
	}

	low := 1
	high := x

	for low < high{
		mid := (low + high)/2
		d := x / mid
		d1 := x/(mid + 1)
		if d >= mid && d1 < mid + 1{
			return mid
		}else if d < mid{
			high = mid-1
		}else{
			low = mid+1
		}
	}

	if low < high {
		return low
	}else{
		return high
	}
}

func main() {
	println(mySqrt(18273))
}


