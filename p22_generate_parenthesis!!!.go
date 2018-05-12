package main

import "fmt"


func isParenthesisValid(str string, n int) bool {
	stack := make([]rune, 0, len(str))
	for _,c := range str{
		if c == '(' {
			stack = append(stack, c)
		}else {
			if len(stack) == 0 {
				return false
			}
			stack = stack[0:len(stack) - 1]
		}
	}
	if len(stack) == 0 {
		return true
	}

	return (n - len(str)) >= len(stack)
}

//n对括号，有n个(，n个)
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	result := make([]string, 0, 100)
	n*=2

	for i:= 0; i<n; i++{

		if len(result) == 0 {
			result = append(result, "(")
			result = append(result, ")")
			continue
		}

		currentLenth := len(result) // len(slice)是实时计算的，每一轮次的长度需预先固定下来
		for j := 0; j < currentLenth; j++ {
			top := result[0]
			result = result[1:]
			candidate1 := top + "("
			candidate2 := top + ")"
			if isParenthesisValid(candidate1, n ) {
				result = append(result, candidate1)
			}

			if isParenthesisValid(candidate2, n) {
				result = append(result, candidate2)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
