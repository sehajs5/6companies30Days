package main

func numFriendRequests(ages []int) int {
	mapp := make(map[int]int)
	for _, age := range ages {
		mapp[age]++
	}
	res := 0

	for key := range mapp {
		for b := range mapp {
			if solve6(key, b) {
				if key == b {
					res += mapp[key] * (mapp[b] - 1)
				} else {
					res += mapp[key] * mapp[b]
				}
			}
		}
	}
	return res
}
func solve6(a, b int) bool {
	return !(b <= a/2+7 || b > a || (b > 100 && a < 100))
}
