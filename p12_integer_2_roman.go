package main

import (
	"math"
	"fmt"
)

func intToRoman(num int) string {
	single := []string{"I","V","IX"}
	ten := []string{"X","L","XC"}
	hundred := []string{"C","D","CM"}
	thousand := []string{"M"}

	bundle := [][]string{single, ten, hundred, thousand}

	var roman string = ""
	for i := 0; i < 4; i++ {
		left := (num/int(math.Pow10(i)))%10
		simbols := bundle[i]
		if left == 0 {
			continue
		}

		if left > 0 && left < 4 {
			for j := 0; j < left; j++ {
				roman = simbols[0] + roman
			}
		} else if (left == 4){
			roman = simbols[0] + simbols[1] + roman
		}  else if left >= 5 && left < 9 {
			for j := 0; j < left-5; j++ {
				roman = simbols[0] + roman
			}
			roman = simbols[1] + roman
		} else {
			roman = simbols[2] + roman
		}
	}
	return roman
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
