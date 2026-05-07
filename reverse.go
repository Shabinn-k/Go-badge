package main

var rev = []int{7, 6, 5, 4, 3, 2, 1}

func Reverse(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
