package main

/**
 n == -2^31 时，如果直接将n取反，则会超过int类型的上界导致溢出
 */
func myPow(x float64, n int) float64 {
	if n >=	 0 {
		return powPositive(x, int64(n))
	}else {
		return  1/powPositive(x, -int64(n))
	}
}

func powPositive(x float64, n int64) float64 {
	if n == 0 {
		return 1
	}
	if x == 0.0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	if n%2 == 0{
		 p := powPositive(x, n/2)
		 return p*p
	}else{
		 p := powPositive(x, (n-1)/2)
		 return x*p*p
	}
}