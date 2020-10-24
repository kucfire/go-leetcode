package leetcode131

func Partition(s string) [][]string {
	result := make([][]string, 0)

	temp := make([]string, 0)
	n := len(s)
	check := func(i, j int) bool {
		for i <= j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			comb := make([]string, len(temp))
			// temp时变动的，需要一个变量来接替temp此时的值
			copy(comb, temp)
			result = append(result, comb)
			// return result
			return
		}

		for i := index; i < n; i++ {
			if !check(index, i) {
				continue
			}

			temp = append(temp, s[index:i+1])
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
		// return result
	}

	dfs(0)
	return result
}
