package main

import "strings"

func isValidSerialization(preorder string) bool {
	// if len(preorder)>=1 && preorder[0]=='#'{
	//     return true
	// }
	// if len(preorder)==1 && preorder[0]!= '#'{
	//     return false
	// }
	// if len(preorder) ==2{
	//     return false
	// }
	places := 1
	for _, n := range strings.Split(preorder, ",") {
		if places > 0 && n == "#" {
			places--
		} else if places > 0 {
			places++
		} else {
			return false
		}
	}
	return places == 0
}
