package main

/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 */

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	sum[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				sum[i][j] = grid[i][j] + sum[i][j-1]
			}else if j == 0 {
				sum[i][j] = grid[i][j] + sum[i-1][j]
			}else{
				sum[i][j] = grid[i][j] + min(sum[i][j-1], sum[i-1][j])
			}
		}
	}

	return sum[m-1][n-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
