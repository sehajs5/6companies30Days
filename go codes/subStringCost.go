package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	l := 0
	curCost := 0
	maxLength := 0

	for r := 0; r < n; r++ {
		curCost += cost(s[r], t[r])
		fmt.Printf("Pass: %d\n", curCost)
		// Shrink the window when the cost exceeds maxCost
		for curCost > maxCost {
			curCost -= cost(s[l], t[l])
			l++
			fmt.Printf("Pass: %d\n", curCost)
		}

		// Update the maximum length of the valid substring
		maxLength = max(maxLength, r-l+1)
		fmt.Printf("Len: %d\n", maxLength)
	}

	return maxLength
}

func cost(s byte, t byte) int {
	// diff := int(s - t)
	// if diff < 0 {
	// 	diff = -diff
	// }
	return int(s) - int(t)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
