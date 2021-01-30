/*
	leetcode tag:139 title:单词拆分
*/

package leetcode139

func WordBreak(s string, wordDict []string) bool {
	// 将wordDict转换成map类型
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	//
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// j代表的是拆分单词的起点，
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
