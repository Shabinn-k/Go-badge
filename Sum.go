package main

var array1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Sum(arr []int) int {
	total := 0
	for _, n := range arr {
		total += n
	}
	return total
}
