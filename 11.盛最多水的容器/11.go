package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for right > left {
		aa := (right - left) * min(height[left], height[right])
		area = max(area, aa)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
