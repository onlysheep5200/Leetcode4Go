package main

import "fmt"

/**
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Input: dividend = 10, divisor = 3
Output: 3

Input: dividend = 7, divisor = -3
Output: -2

Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 2^31 − 1 when the division result overflows.

 */

const INT_MAX int = int(^uint32(0) >> 1)
const INT_MIN int = ^INT_MAX


 //除法的本意，即被除数一次次减去除数，求  得到最小非负结果的 次数
 //还需考虑越界的情况（INT_MIN的绝对值比INT_MAX多1）
 //golang的正数类型还是需要注意位数滴
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return dividend
	}

	//越界的情况
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	dividendLong := int64(dividend)
	if dividendLong <0 {
		dividendLong = -dividendLong
	}

	divisorLong := int64(divisor)
	if divisorLong < 0 {
		divisorLong = -divisorLong
	}

	if divisorLong > dividendLong {
		return 0
	}

	symbol := 1
	if divisor > 0 && dividend < 0 || divisor < 0 && dividend > 0 {
		symbol = -1
	}
	divisorMoved := divisorLong
	var moveCount uint32 = 0
	count := 0

	for dividendLong >= divisorMoved {
		dividendLong -= divisorMoved
		divisorMoved = divisorMoved << 1
		count += 1<< moveCount
		moveCount++
		if dividendLong < divisorMoved {
			divisorMoved = divisorLong
			moveCount = 0
		}
	}

	return symbol*count
}

func main() {
	fmt.Println(divide(INT_MIN, -2))
}