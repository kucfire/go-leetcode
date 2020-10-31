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
	for head := 0; head < len(s); head++ {
		tmp := s[:head] + s[head+1:]
	loop2:
		for tail := 0; tail < len(tmp); tail++ {
			if s[head] == tmp[tail] {
				break loop2
			}
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
