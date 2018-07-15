package main

import "fmt"

/**
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like ? or *.
 */

func isMatch(s string, p string) bool {
	//return matchPartial(s, p)
	return isMatchFSM(s, p)
}

/**
有限状态机法
 */
func isMatchFSM(s string, p string) bool{

	type line struct {
		state int
		token uint8
	}

	initState := 0

	buildFSM:= func() (map[line]int, int){
		state := initState
		fsm := make(map[line]int)
		for i:=0; i<len(p); i++{
			chr := p[i]
			if chr != '*'{
				fsm[line{state:state, token:chr}] = state+1
				state++
			}else{
				fsm[line{state:state, token:chr}] = state
			}
		}
		return fsm, state
	}

	fsm, finishState := buildFSM()
	stateSet := make(map[int]bool)
	stateSet[0] = true

	for i := 0; i < len(s); i++ {
		chr := s[i]
		newStateSet := make(map[int]bool)
		for state := range stateSet {
			nextState,exists := fsm[line{state: state, token: chr}]
			if exists{
				newStateSet[nextState] = true
			}

			nextState,exists = fsm[line{state: state, token: '*'}]
			if exists{
				newStateSet[nextState] = true
			}

			nextState,exists = fsm[line{state: state, token: '?'}]
			if exists{
				newStateSet[nextState] = true
			}
		}
		//每执行一步，需刷新状态集合
		stateSet = newStateSet
	}

	return stateSet[finishState]
}

/** TLE 普通的递归**/
func matchPartial(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) == 0 && len(p) != 0 {
		for _,chr := range p{
			if chr != '*' {
				return false
			}
		}
	}
	sIndex := 0
	pIndex := 0

	for sIndex < len(s) && pIndex < len(p){
		switch p[pIndex] {
		case '?':
			sIndex++
			pIndex++
			continue
		case '*':
			if pIndex == len(p) - 1 {
				return true
			}
			//找到第一个部位*的字符
			var nextRune uint8
			nextPIndex := pIndex
			for nextRune = p[nextPIndex]; nextPIndex < len(p); nextPIndex++{
				nextRune = p[nextPIndex]
				if nextRune != '*' {
					break
				}
			}
			if nextRune == '*' {
				return true
			}

			nextSIndex := findNextSIndex(s, sIndex, nextRune)
			leftMatch := false
			if nextSIndex >= len(s) {
				return false
			}
			for !leftMatch {
				leftMatch = leftMatch ||  matchPartial(s[nextSIndex : ], p[nextPIndex : ])
				if leftMatch {
					return true
				}
				nextSIndex = findNextSIndex(s, nextSIndex+1, nextRune)
				if nextSIndex >= len(s){
					return false
				}
			}
			return leftMatch
		default:
			if s[sIndex] != p[pIndex] {
				return false
			}
			sIndex++
			pIndex++
			continue
		}
	}

	if sIndex < len(s) {
		return false
	}

	if pIndex < len(p) {
		for ; pIndex < len(p); pIndex++ {
			if p[pIndex] != '*' {
				return false
			}
		}
	}

	return true
}


func findNextSIndex(s string, i int, u uint8) int {
	if i > len(s) {
		return len(s)
	}
	if u == '?' {
		return i
	}

	for j := i; j<len(s); j++{
		if s[j] == u {
			return j
		}
	}
	return len(s)
}

func main() {
	fmt.Println(isMatch("",""))
}