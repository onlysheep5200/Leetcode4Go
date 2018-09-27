package main

import "fmt"

/**
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character
Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

 */

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	d := make([][]int, 0, m+1)

	for i := 0; i <= m+1; i++ {
		d = append(d, make([]int, n+1))
	}

	for i := 0; i < n+1; i++ {
		d[0][i] = i
	}

	for j := 0; j < m+1; j++ {
		d[j][0] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1]{
				d[i][j] = d[i-1][j-1]
				continue
			}
			d[i][j] = min3(d[i-1][j-1] + 1, d[i-1][j] +1, d[i][j-1] + 1)
		}
	}

	return d[m][n]
}

func min3(a, b, c int) int {
	if a > b{
		a = b
	}

	if a > c {
		a = c
	}

	return a
}

func main() {
	fmt.Println(minDistance("intention", "execution"))
}
