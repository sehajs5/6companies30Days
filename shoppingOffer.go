package main

func shoppingOffers(price []int, special [][]int, needs []int) int {
	pricess := solve2(price, special, needs, 0)
	return pricess
}

func solve2(price []int, special [][]int, needs []int, pos int) int {
	n := len(needs)
	currentPrice := 0
	for i := 0; i < n; i++ {
		currentPrice += (needs[i] * price[i])
	}
	for i := pos; i < len(special); i++ {
		offer := special[i]
		arr := []int{}
		for j := 0; j < len(needs); j++ {
			if needs[j] < offer[j] {
				arr = nil
				break
			}
			arr = append(arr, needs[j]-offer[j])
		}

		if len(arr) > 0 {
			currentPrice = min(currentPrice, offer[n]+(solve2(price, special, arr, i)))
		}
	}
	return currentPrice
}
