package main

func longestPalindrome(s string) string {
	maxS := ""
	for i := 0; i < len(s); i++ {
		s1, s2 := getMaxLen(s, i, i), getMaxLen(s, i, i+1)
		if len(s1) > len(maxS) {
			maxS = s1
		}
		if len(s2) > len(maxS) {
			maxS = s2
		}
	}
	return maxS
}

func getMaxLen(s string, i, j int) string {
	l, r := 0, len(s)-1
	for i >= l && j <= r {
		if s[i] == s[j] {
			i--
			j++
		} else {
			break
		}
	}

	return s[i+1 : j]
}
