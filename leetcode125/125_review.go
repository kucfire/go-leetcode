package leetcode125

import "strings"

func IsPalindromeReview(s string) bool {
	s = strings.ToLower(s)
	IsAlnum := func(str byte) bool {
		return (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') || (str >= '0' && str <= '9')
		//  (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') || (str >= '0' && str <= '9')
	}

	start, end := 0, len(s)-1
	for start < end {
		for !IsAlnum(s[start]) && start < end {
			start++
		}
		for !IsAlnum(s[end]) && start < end {
			end--
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
