/*
	leetcode tag:461 title:汉明距离
*/

package leetcode461

import "fmt"

// HammingDistance : bit operation
func HammingDistance(x int, y int) int {
	var result int

	// fmt.Println(i)
	for i := x ^ y; i != 0; i = i >> 1 {
		if i&1 == 1 {
			fmt.Println(i, i&1)
			result++
		}
	}

	return result
}
