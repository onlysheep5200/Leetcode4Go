package main

import "fmt"

/**
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?


Above is a 7 x 3 grid. How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:

Input: m = 7, n = 3
Output: 28
 */

func uniquePaths(m int, n int) int {
	area := make([][]int, m)
	for i := 0; i < m; i++ {
		area[i] =  make([]int, n)
	}

	area[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			if i > 0 {
				count += area[i-1][j]
			}

			if j > 0 {
				count += area[i][j-1]
			}
			if count > 0 {
				area[i][j] = count
			}
		}
	}

	return area[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(7, 3))
}