package main

import (
	"strconv"
	"fmt"
)

// FIFO
// 有序，组合，广度优先
func letterCombinations(digits string) []string {
	numLetterMap := make(map[int64][]string)
	numLetterMap[2] = []string{"a", "b", "c"}
	numLetterMap[3] = []string{"d", "e", "f"}
	numLetterMap[4] = []string{"g", "h", "i"}
	numLetterMap[5] = []string{"j", "k", "l"}
	numLetterMap[6] = []string{"m", "n", "o"}
	numLetterMap[7] = []string{"p", "q", "r", "s"}
	numLetterMap[8] = []string{"t", "u", "v"}
	numLetterMap[9] = []string{"w", "x", "y", "z"}

	result := make([]string, 0, 100)

	for _,char := range digits{
		i, _ := strconv.ParseInt(string(char), 10, 0)
		candidates := numLetterMap[i]
		if len(result) == 0 {
			result = append(result, candidates...)
		}else{
			currentLen := len(result)
			for j := 0; j < currentLen; j++ {
				top := result[0]
				for _,candidate := range candidates{
					result = append(result, top + candidate)
				}
				result = result[1:]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
