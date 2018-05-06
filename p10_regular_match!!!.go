package main

import "golang.org/x/tools/go/ssa"

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		if len(s) != 0 {
			return false
		}else{
			return true;
		}
	}

	if len(p) == 1 {
		if len(s) != 1{
			return false
		}

	}
}

func main() {

}
