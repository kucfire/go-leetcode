/*
	leetcode tag:022 title:括号生成
*/

package leetcode022

// GenerateParenthesis : method of recursion
func GenerateParenthesis(n int) []string {
	var result []string
	parenthesis := func(l, r int, result *[]string, p string) {}
	parenthesis = func(l, r int, result *[]string, p string) {
		// 如果l,r都为0，则表明进入了最底一层，开始进行append
		if l == 0 && r == 0 {
			*result = append(*result, p)
		}

		// 剩余的左括号不能大于右括号
		if l > r || l < 0 || r < 0 {
			return
		}
		parenthesis(l-1, r, result, p+"(")
		parenthesis(l, r-1, result, p+")")
	}
	parenthesis(n, n, &result, "")
	return result
}

// method of recall 回溯
func GenerateParenthesis2(n int) []string {
	var result []string
	var parenthesis func(l, r int, path string)
	parenthesis = func(l, r int, path string) {
		// 当长度达到两倍n的时候，将其添加进结果里面
		if 2*n == len(path) {
			result = append(result, path)
			return
		}

		if l > 0 {
			parenthesis(l-1, r, path+"(")
		}
		if r > 0 && r > l {
			parenthesis(l, r-1, path+")")
		}
	}
	parenthesis(n, n, "")
	return result
}
