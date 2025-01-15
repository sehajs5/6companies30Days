package main

func firstUniqChar(s string) int {
	mapp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapp[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if mapp[s[i]] == 1 {
			return i
		}
	}
	return -1
}
