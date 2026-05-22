package main

var nums = []int{2, 7, 11, 15}
var target = 9

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if j, found := m[need]; found {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
