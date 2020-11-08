/*
	leetcode tag:121 title:买卖股票的最佳时机
*/

package leetcode121

// MaxProfit : one time to get the most money(method of double loop)
func MaxProfit(prices []int) int {
	// 排除异常传入数值
	if len(prices) < 2 {
		return 0
	}
	var result int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			tmp := prices[j] - prices[i]
			if tmp > 0 && tmp > result {
				result = tmp
			} else {

			}
		}
	}
	return result
}

// MaxProfit2 :
func MaxProfit2(prices []int) int {
	// 排除异常传入数值
	if len(prices) < 2 {
		return 0
	}
	result, head, tail, tmp, length := 0, 0, 1, 0, len(prices)
	for head < tail && head < length-1 && tail < length {
		tmp = prices[tail] - prices[head]
		if tmp < 0 { // 代表有耕地的买入价，head可以看作是寻找len(prices)-1中最低的买入价格，tail可以看做是寻找1~len(prices)中最高的卖出价格，result记录最大的利润值
			head = tail
			tail++
			continue
		}
		if tmp > result {
			result = tmp
		}
		tail++
	}
	return result
}
