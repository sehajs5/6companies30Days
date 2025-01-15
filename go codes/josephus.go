package main

func findTheWinner(n int, k int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans += k
		ans = (ans-1)%i + 1
	}
	return ans
}
