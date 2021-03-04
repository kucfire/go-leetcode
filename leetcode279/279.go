/*
	leetcode tag:279 title:完全平方数
*/

package leetcode279

import "math"

// 抄的
func NumSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		x := int(math.Sqrt(float64(i)))
		if x*x == i {
			dp[i] = 1
		} else {
			dp[i] = n
			for k := x; k > 0; k-- {
				if dp[i] > dp[i-k*k]+1 {
					dp[i] = dp[i-k*k] + 1
				}
			}
		}
	}
	return dp[n]
}
