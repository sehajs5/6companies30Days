package main

const MOD = 1e9 + 7

func squareFreeSubsets(nums []int) int {
	dp := make([]int, 1<<10)
	dp[0] = 1
	for _, num := range nums {
		mask := getMask(num)
		if mask < 0 {
			continue
		}

		for i := 0; i < (1 << 10); i++ {
			if (mask & i) == 0 {
				newProduct := mask | i
				dp[newProduct] += (dp[i] % MOD)
			}
		}
	}
	sum := 0
	for _, val := range dp {
		sum += val
	}
	return (sum - 1) % MOD
}

func getMask(num int) int {
	var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	mask := 0

	for i, prime := range primes {
		if num%(prime*prime) == 0 {
			mask = -1
			break
		} else if num%prime == 0 {
			mask = mask | (1 << i)
		}
	}
	return mask
}
