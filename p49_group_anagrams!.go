package main

import (
	"strings"
	"fmt"
)

/**
参考提交中耗时较少的答案的写法
 */
func groupAnagrams(strs []string) [][]string {
	result := make([][]string,0)
	if len(strs) == 0 {
		return result
	}

	wordMap := make(map[string]int)

	for _,str := range strs{
		sortedStr := sortString(str)
		index,exists := wordMap[sortedStr]

		if exists {
			result[index] = append(result[index], str)
		}else{
			wordMap[sortedStr] = len(result)
			result = append(result, []string{str})
		}
	}

	return result
}

/**
将字符串根据各个字母的字段序排序
 */
func sortString(str string)string  {
	if len(str) == 0 {
		return ""
	}

	//此处只可能为26个小写字母
	charNums := make([]int, 26)

	for _,c := range str{
		charNums[c-'a']++
	}

	var ns strings.Builder
	for i := 0; i < 26; i++{
		if charNums[i] == 0 {
			continue
		}
		for j := 0; j < charNums[i]; j++ {
			ns.WriteRune(rune('a' + i))
		}
	}

	return ns.String()
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
