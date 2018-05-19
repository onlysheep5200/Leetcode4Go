package main

import "fmt"

/**
You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without
any intervening characters.
 */


 //备选的字符串等长，可将主串截取一段，等分后判断各个子串是否能在候选的words中找到
func findSubstring(s string, words []string) []int {
	result := make([]int, 0, 100)
	if len(s) == 0 || len(words) == 0 {
		return result
	}

	subStrLength := len(words[0]) * len(words)
	wordLength := len(words[0])
	if subStrLength > len(s) {
		return result
	}

	wordsMap := make(map[string]int)

	for _,word := range words{
		wordsMap[word] += 1
	}

	for i := 0; i < len(s)-subStrLength+1; i++ {
		subStr := s[i : i+subStrLength]
		candidateMap := make(map[string]int)
		wordCount := 0
		for j := 0; j< subStrLength-wordLength+1; j+=wordLength  {
			w := subStr[j : j+wordLength]
			if count, ok := wordsMap[w]; ok {
				if candidateMap[w] < count {
					candidateMap[w] += 1
					wordCount++
				}else{
					break
				}
			}
		}
		if wordCount == len(words){
			result = append(result, i)
		}
	}


	//golang的slice作为函数参数，如果有append等操作，需要传指针
	//tranverse("", subStrLength, &words, candidates, &used, &result)   tle
	return result
}

//func tranverse(current string, substrLength int, words *[]string, candidates map[string][]int, used *[]bool, result *[]int){
//	if len(current) == substrLength {
//		if candidate,ok := candidates[current]; ok {
//			*result = append(*result, candidate...)
//			delete(candidates, current)
//		}
//		return
//	}
//	for i := 0; i < len(*words); i++ {
//		if !(*used)[i] {
//			(*used)[i] = true
//			tranverse(current+(*words)[i], substrLength, words, candidates, used, result)
//			(*used)[i] = false
//		}
//	}
//}

func main() {
	 s := "barfoothefoobarman";
	 words := []string{"foo","bar"}

	 fmt.Println(findSubstring(s, words))

}