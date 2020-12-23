/*
	leetcode tag:387 title:字符串中的第一个唯一字符
*/

package leetcode387

// FirstUniqChar : violence method
func FirstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}
	// head := 0
	// 双重循环
	// 第一重循环：取出对比的子字符串，并将剩余的字符串放入临时变量tmp内
	for head := 0; head < len(s); head++ {
		tmp := s[:head] + s[head+1:]
		// 第二重循环，遍历tmp，如果发生和第一重循环取出来的子字符串相等的话，立马跳出第二重循环，并重新进去第一重循环
	loop2:
		for tail := 0; tail < len(tmp); tail++ {
			if s[head] == tmp[tail] {
				break loop2
			}
			// 如果遍历到最后还是没找到与head相对应的子字符串相等的元素话，直接返回该head
			if tail == len(tmp)-1 {
				return head
			}
		}
	}
	return -1
}

// FirstUniqChar2 : method of array
func FirstUniqChar2(s string) int {
	arrayList := make([]int, 26)
	for i, str := range s {
		arrayList[str-'a'] = i
	}
	for i, str := range s {
		if i == arrayList[str-'a'] {
			return i
		} else {
			arrayList[str-'a'] = -1
		}
	}
	return -1
}
