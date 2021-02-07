/*
	leetcode tag:334 title:递增的三元子序列
*/

package leetcode334

import "math"

func IncreasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if m1 >= num {
			m1 = num
		} else if m2 >= num {
			m2 = num
		} else {
			return true
		}
	}

	return false
}
