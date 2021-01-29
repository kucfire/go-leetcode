package leetcode887

// 动态规划
// func SuperEggDropReview(k, n int) int {
// 	result := make([][]int, k+1)
// 	for i := range result {
// 		result[i] = make([]int, n+1)
// 	}

// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= k; j++ {

// 			result[j][i] = result[j][i-1] + result[j-1][i-1] + 1
// 			if result[j][i] >= n {
// 				return i
// 			}
// 		}
// 	}

// 	return n
// }

func SuperEggDropReview(k, n int) int {
	result := make([][]int, n+1)
	for i := range result {
		result[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j] + 1
			if result[i][j] >= n {
				return i
			}
		}
	}

	return n
}
