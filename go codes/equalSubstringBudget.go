package main

func equalSubstring2(s string, t string, maxCost int) int {
	j := 0
	maxi := 0
	curCost := 0
	for i := 0; i < len(s); i++ {
		diff := int(s[i]) - int(t[i])
		if diff < 0 {
			diff = -(diff)
		}
		curCost += diff

		for curCost > maxCost {
			diff = int(s[j]) - int(t[j])
			if diff < 0 {
				diff = (-diff)
			}
			curCost -= diff
			j++
		}
		maxi = max(maxi, i-j+1)
	}
	return maxi
}
