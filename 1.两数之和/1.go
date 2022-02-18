package main

func twoSum(nums []int, target int) []int {
	num := make(map[int]int)
	for k, v := range nums {
		if kk, ok := num[target-v]; ok {
			return []int{kk, k}
		}
		num[v] = k
	}
	return nil
}
