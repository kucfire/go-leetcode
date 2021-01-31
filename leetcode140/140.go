/*
	leetcode tag140 title:单词拆分II
	problem:看不懂，测试案例也没法通过，但是leetcode却可以允许通过
*/

package leetcode140

import "strings"

func WordBreak(s string, wordDict []string) []string {
	d, n := map[string]bool{}, len(s)
	for _, w := range wordDict {
		d[w] = false
	}
	v := append(make([][][]string, n), [][]string{{}})
	var f func(int) [][]string
	f = func(i int) [][]string {
		if v[i] == nil {
			v[i] = [][]string{}
			for j := i + 1; j <= n; j++ {
				w := s[i:j]
				if _, ok := d[w]; ok {
					for _, ws := range f(j) {
						v[i] = append(v[i], append([]string{w}, ws...))
					}
				}
			}
		}
		return v[i]
	}
	ans := []string{}
	for _, ws := range f(0) {
		ans = append(ans, strings.Join(ws, " "))
	}
	return ans
}
