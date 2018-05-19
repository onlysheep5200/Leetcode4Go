package main

import "fmt"

/**
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
 */

 //栈里面放index不就好了
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxValidLen := 0
	validLen := 0
	stack := make([]int, 0 , len(s) + 1)
	stack = append(stack, -1)

	for index,chr := range s {
		if chr == '(' {
			stack = append(stack, index)
		}else if len(stack) == 1 {
			//除了-1啥也没有，说明 ) 多了
			//此时开始未知应移到当前 ) 所在的位置
			stack[0] = index
			continue
		}else {
			stack = stack[0:len(stack) - 1]
			peek := stack[len(stack) - 1]
			validLen = index - peek
			if validLen > maxValidLen {
				maxValidLen = validLen
			}
		}
	}

	return maxValidLen
}

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}