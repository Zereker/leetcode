package main

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return fastPow(x, n)
	}

	return 1.0 / fastPow(x, -n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	pow := fastPow(x, n/2)
	if n%2 == 0 {
		return pow * pow
	}

	return pow * pow * x
}
