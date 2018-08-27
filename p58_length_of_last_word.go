package main

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5

 */


func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	nonBlankOccur := false
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if nonBlankOccur{
				break
			}else{
				continue
			}
		}
		nonBlankOccur = true
		count++
	}

	return count
}