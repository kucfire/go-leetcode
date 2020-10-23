/*
	leetcode tag:125 title:验证回文串
*/

package leetcode125

import "strings"

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	isalnum := func(str byte) bool {
		return (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') || (str >= '0' && str <= '9')
	}
	left, right := 0, len(s)-1
	for right > left {
		for !isalnum(s[left]) && right > left {
			left++
		}
		for !isalnum(s[right]) && right > left {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
