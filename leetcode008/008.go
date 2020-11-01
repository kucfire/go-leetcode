/*
	leetcode tag:008 title:字符串转换整数(atoi)
*/

package leetcode008

import (
	"math"
	"strings"
)

// MyAtoi : normal method
func MyAtoi(s string) int {
	isUseful := func(s string) int {
		for i, str := range s {
			if str == ' ' {
				continue
			}
			if str == '-' || str == '+' || (str >= '0' && str <= '9') {
				return i
			} else {
				return -1
			}
		}
		return -1
	}

	usefulIndex := isUseful(s)
	if usefulIndex == -1 {
		return 0
	}

	result := 0
	isMinus := false
	for i := usefulIndex; i < len(s); i++ {
		// 判断第一个符号是不是负号
		if i == usefulIndex && s[i] == '-' {
			isMinus = true
			continue
		} else if i == usefulIndex && s[i] == '+' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i]-'0')
			// log.Println(result)
			if result > math.MaxInt32 {
				if isMinus == true {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		} else {
			break
		}
	}
	if isMinus == true {
		result = -result
	}
	return result
}

// MyAtoi2 : the best method
func MyAtoi2(s string) int {
	// clear all the sapce " "
	s = strings.TrimSpace(s)
	result := 0
	sign := 0
	for i, str := range s {
		if i == 0 && str == '-' {
			sign = 1
			continue
		} else if i == 0 && str == '+' {
			sign = 0
			continue
		}
		if str >= '0' && str <= '9' {
			result = result*10 + int(str-'0')
		} else {
			break
		}
		if result > math.MaxInt32 {
			if sign == 1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if sign == 1 {
		result = -result
	}
	return result
}
