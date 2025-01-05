package main

var ans int

func maxProduct(s string) int {
	ans = 0
	solve1(s, 0, "", "")
	return ans
}
func solve1(s string, i int, s1 string, s2 string) {
	if isPal(s1) && isPal(s2) {
		ans = max(ans, len(s1)*len(s2))
	}

	if i == len(s) {
		return
	}
	solve1(s, i+1, s1+string(s[i]), s2)

	solve1(s, i+1, s1, s2+string(s[i]))

	solve1(s, i+1, s1, s2)

	return
}

func isPal(s string) bool {
	n := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
