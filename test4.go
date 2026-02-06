package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 {
		if divisibleBy(n, 2) {
			n /= 2
		} else if divisibleBy(n, 3) {
			n /= 3
		} else if divisibleBy(n, 5) {
			n /= 5
		} else {
			return false
		}
	}
	return true
}

func divisibleBy(n int, div int) bool {
	return n%div == 0
}
