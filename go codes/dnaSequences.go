package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}

	ans := []string{}
	mapp := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		sample := string(s[i : i+10])

		if mapp[sample] == 1 {
			ans = append(ans, sample)
		}
		mapp[sample]++

	}
	return ans
}
