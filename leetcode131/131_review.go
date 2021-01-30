package leetcode131

func PartitionReview(s string) [][]string {
	// 定义变量
	result := make([][]string, 0) // 返回结果
	temp := make([]string, 0)     // 作为添加进result的临时一元数组
	sLen := len(s)                // 字符串s的长度

	// 检查是否为回文字符串
	check := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	dps := func(index int) {}
	dps = func(index int) {
		if index == sLen {
			app := make([]string, len(temp))
			copy(app, temp)
			result = append(result, app)
			return // 动态规划最底层的返回
		}

		for i := index; i < sLen; i++ {
			if !check(index, i) {
				continue
			}
			temp = append(temp, s[index:i+1])
			dps(i + 1)
			temp = temp[:len(temp)-1] // *
		}
	}

	dps(0)
	return result
}
