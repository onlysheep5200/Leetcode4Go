package main


func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	common := strs[0]
	if common == "" {
		return common
	}

	for i := 1; i < len(strs); i++ {
		curStr := strs[i]
		if curStr == "" {
			 return ""
		}
		size := len(common)
		if len(curStr) < size {
			size = len(curStr)
		}
		if common[0] != curStr[0] {
			return ""
		}

		for j := 1; j < size; j++ {
			if common[j] != curStr[j] {
				common = common[:j]
				break
			}
		}

		if len(common) > size {
			common = common[:size]
		}
	}

	return common
}

