/*
	leetcode tag:122 title:买卖股票的最佳时机II
*/

package leetcode122

// MaxProfit : method of normal
func MaxProfit(prices []int) int {
	result := 0
	for i, num := range prices {
		if i == 0 {
			continue
		}
		if num > prices[i-1] {
			result += num - prices[i-1]
		}
	}
	return result
}
