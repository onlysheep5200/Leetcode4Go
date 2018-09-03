package main

import (
	"strings"
	"fmt"
)

/**
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

Corner Cases:

Did you consider the case where path = "/../"?
In this case, you should return "/".
Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
In this case, you should ignore redundant slashes and return "/home/foo".
 */

func modifyStack(stack []string, sub string) []string {
	if sub == "."{
		return stack
	}
	if sub == ".." && len(stack) != 0{
		stack = stack[0:len(stack) - 1]
	}else if sub != ".." {
		stack = append(stack, sub)
	}

	return stack
}

func simplifyPath(path string) string {
	if len(path) == 0 {
		return path
	}
	stack := make([]string, 0, len(path))

	sub := ""
	for _,c := range path{
		if c == '/' &&  sub != ""{
			stack = modifyStack(stack, sub)
			sub = ""
		}else if c != '/'{
			sub = fmt.Sprintf("%s%c", sub, c)
		}
	}

	//可能还有残留，此处需要处理
	if sub != ""{
		stack = modifyStack(stack, sub)
	}
	path = "/" + strings.Join(stack, "/")
	return path
}

func main() {
	println(simplifyPath("/..."))
}
