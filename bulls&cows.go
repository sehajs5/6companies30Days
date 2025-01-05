package main

import "fmt"

func getHint(secret string, guess string) string {
	bulls := getbulls(secret, guess)
	cows := getcows(secret, guess)
	cows = cows - bulls
	ans := fmt.Sprintf("%d%c%d%c", bulls, 'A', cows, 'B')
	return ans
}
func getcows(secret string, guess string) int {
	mapp := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		mapp[secret[i]]++
	}
	ans := 0
	for i := 0; i < len(guess); i++ {
		if val, found := mapp[guess[i]]; found && val > 0 {
			ans++
			mapp[guess[i]]--
		}
	}
	return ans
}
func getbulls(secret string, guess string) int {
	ans := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			ans++
		}
	}
	return ans
}
