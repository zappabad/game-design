package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	var digits []int = []int{}
	if x < 0 {
		return false
	}
	num_of_digits := int(math.Log10(float64(x))) + 1
	for i := 0; i < num_of_digits; i++ {
		digit := extractDigitAtIndex(x, i)
		digits = append(digits, digit)
	}
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

func extractDigitAtIndex(x int, index int) int {
	modulo_by := int(math.Pow(10, float64(index+1)))
	div_by := int(math.Pow(10, float64(index)))
	return (x % modulo_by) / div_by
}

func makeNumberFromDigits(digits []int) int {
	var number int = 0
	for i := 0; i < len(digits); i++ {
		power := int(math.Pow(10, float64(i)))
		number += digits[i] * power
	}
	return number
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(12123121))
	fmt.Println(isPalindrome(125341))
	fmt.Println(isPalindrome(12199121))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(-121))
}
