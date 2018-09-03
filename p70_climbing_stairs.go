package main

/**
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 */

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	p := make([]int, n)
	p[0] = 1
	p[1] = 1
	for i := 1; i < n; i++ {
		p[i] += p[i-1]
		if i-2 >= 0 {
			p[i] += p[i-2]
		}
	}

	return p[n-1]
}

func main() {
	println(climbStairs(3))
}