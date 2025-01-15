package main

func convertToTitle(columnNumber int) string {
	if columnNumber < 27 {
		return string(rune(columnNumber + 64))
	}
	ans := ""
	for columnNumber > 0 {
		columnNumber--
		ans = string(rune(columnNumber%26+'A')) + ans
		columnNumber /= 26
	}
	return string(ans)

}
