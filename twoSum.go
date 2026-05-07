package main

var nums = []int{2, 7, 11, 15}
var target = 9

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		get := target - num
		if ind, found := m[get]; found {
			return []int{ind, i}
		}
		m[num] = i
	}
	return nil
}
