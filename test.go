package main

import (
	"fmt"
	"math"
)

func canMakeArithmeticProgression(arr []int) bool {
	sorted_arr := sortListMintoMax(arr)
	var diffs []int = []int{}
	for i := 0; i < len(sorted_arr)-1; i++ {
		var j int = i + 1
		diff := diffElements(sorted_arr[i], sorted_arr[j])
		diffs = append(diffs, diff)
	}
	return checkIfAllEqual(diffs)
}

func diffElements(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func checkIfAllEqual(diffs []int) bool {
	for i := 1; i < len(diffs); i++ {
		if diffs[i] != diffs[0] {
			return false
		}
	}
	return true
}

func sortListMintoMax(arr []int) []int {
	// sorts the array in ascending order, returning a new array, not modifying the original one
	sorted_arr := make([]int, len(arr))
	copy(sorted_arr, arr)
	for i := 0; i < len(sorted_arr)-1; i++ {
		for j := 0; j < len(sorted_arr)-i-1; j++ {
			if sorted_arr[j] > sorted_arr[j+1] {
				sorted_arr[j], sorted_arr[j+1] = sorted_arr[j+1], sorted_arr[j]
			}
		}
	}
	fmt.Println(sorted_arr)
	return sorted_arr
}

func main() {
	var arr []int = []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}
