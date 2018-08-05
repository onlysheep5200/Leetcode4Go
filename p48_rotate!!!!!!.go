package main

/**yan
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
 */

func rotate(matrix [][]int)  {
	//顺时针先转，先上下翻折，再沿着主对角线翻折
	if len(matrix) == 0 || len(matrix[0]) == 0 || len(matrix) != len(matrix[0]) {
		return
	}

	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[j][i], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[j][i]
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
