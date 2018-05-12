package main

import "fmt"

func isValid(s string) bool {
	pMap := make(map[rune]rune)
	pMap['}'] = '{'
	pMap[')'] = '('
	pMap[']'] = '['

	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0, len(s))
	for _,c := range s{
		if mapping,exist := pMap[c]; exist{
			if len(stack) == 0{
				return false
			}
			peek := stack[len(stack) - 1]
			stack = stack[0:len(stack) - 1]
			if peek != mapping {
				return false
			}
		}else{
			stack = append(stack, c)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}