package main

import "fmt"

func maxProfit(k int, prices []int) int {
	memo := make(map[string]int)
	return solve3(prices, memo, false, 0, k, 0)
}
func solve3(prices []int, memo map[string]int, buy bool, ind int, i int, k int) int {
	if ind == len(prices) || k == i {
		return 0
	}
	key := fmt.Sprintf("%d:%t:%d", ind, buy, k)
	if val, ok := memo[key]; ok {
		return val
	}
	maxProfit := 0
	if !buy {
		one := -prices[ind] + solve3(prices, memo, true, ind+1, i, k)
		two := solve3(prices, memo, false, ind+1, i, k)
		maxProfit = max(one, two)
	} else {
		oneL := prices[ind] + solve3(prices, memo, false, ind+1, i, k+1)
		twoL := solve3(prices, memo, true, ind+1, i, k)
		maxProfit = max(oneL, twoL)
	}
	memo[key] = maxProfit
	return maxProfit
}
