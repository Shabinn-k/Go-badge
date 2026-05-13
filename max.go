package main

var maxNums = []int{1, 2, 3, 4, 6, 7, 8, 9}

func Max(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}else{
			min = arr[i]
		}
	}
	return max, min
}

