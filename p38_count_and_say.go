package main

import "fmt"

/**
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.
 */

 //字符串转数字集合操作会更快一些
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	start := "1"
	count := 0
	var prev byte = 0
	var result string



	for i := 1; i < n; i++ {
		result = ""
		prev = 0
		count = 0
		for j := 0; j<len(start); j++{
			if prev == 0 {
				count = 1
				prev = start[j]
			}else{
				if prev == start[j] {
					count ++
				}else{
					result = fmt.Sprintf("%s%d%c", result, count, prev)
					prev = start[j]
					count = 1
				}
			}
		}

		if prev != 0 {
			result = fmt.Sprintf("%s%d%c", result, count, prev)
		}

		start = result
	}

	return result
}

func main() {
	fmt.Println(countAndSay(5))
}