package main

import "fmt"

/**
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"
 */

func getPermutation(n int, k int) string {
	visited := make([]bool, n+1)
	totalCount := 1
	for i := n-1; i >= 1; i-- {
		totalCount *= i
	}
	result := ""

	for i:=0; i< n; i++{
		t := k/totalCount
		mod := k%totalCount
		if mod == 0 {
			t-=1
		}
		ot := t
		cur := -1
		for m := 1; m < n+1; m++ {
			if visited[m] {
				continue
			}
			cur = m
			t--
			if t < 0 {
				break
			}
		}
		visited[cur] = true
		result = fmt.Sprintf("%s%d", result, cur)
		k -= totalCount*ot
		if totalCount > 1{
			totalCount = totalCount/(n-i-1)
		}
	}

	return result
}

func getPermutations(prefix string, visited []bool, container []string)[]string {
	if len(prefix) == len(visited)-1 {
		container = append(container, prefix)
		return container
	}

	for i := 1; i < len(visited); i++ {
		if !(visited[i]) {
			visited[i] = true
			container = getPermutations(fmt.Sprintf("%s%d", prefix, i), visited, container)
			visited[i] = false
		}
	}

	return container
}

func main() {
	fmt.Println(getPermutation(2,1))
}