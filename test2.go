package main

import "fmt"

func pivotInteger(n int) int {
	list_1_to_n := makeSeqOfIntegers(1, n)
	for i := 1; i <= n; i++ {
		list_1_to_i, list_i_to_n := splitListInIndex(list_1_to_n, i)

		sum_of_pivot_1 := sumUpToIndex(list_1_to_i, len(list_1_to_i))
		sum_of_pivot_2 := sumUpToIndex(list_i_to_n, len(list_i_to_n))
		if sum_of_pivot_1 == sum_of_pivot_2 {
			return i
		}
	}
	return -1
}

func splitListInIndex(list_of_ints []int, index int) ([]int, []int) {
	return list_of_ints[:index], list_of_ints[index-1:]
}

func sumUpToIndex(list_of_ints []int, index int) int {
	var sum int = 0
	for i := 0; i < index && i < len(list_of_ints); i++ {
		sum += list_of_ints[i]
	}
	return sum
}

func makeSeqOfIntegers(start_int int, end_int int) []int {
	list_of_ints := []int{}
	for i := start_int; i <= end_int; i++ {
		list_of_ints = append(list_of_ints, i)
	}
	return list_of_ints
}

func main() {
	var n int = 8
	fmt.Println(pivotInteger(n))
}
