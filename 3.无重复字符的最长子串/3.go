package main

func lengthOfLongestSubstring(s string) int {
	left, res, n := -1, 0, len(s)
	m := make(map[byte]int)
	for right := 0; right < n; right++ {
		ch := s[right]
		m[ch]++
		for m[ch] > 1 {
			left++
			m[s[left]]--
		}
		res = max(res, right-left)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
