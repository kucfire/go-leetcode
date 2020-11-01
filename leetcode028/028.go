/*
	leetcode tag:028 title:实现str
*/

package leetcode028

import "strings"

// StrStr : normal method
func StrStr(haystack string, needle string) int {
	tmp := -1

	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	needleLength := len(needle)
	for i, str := range haystack {
		if byte(str) == needle[0] {
			if i+needleLength <= len(haystack) {
				if string(haystack[i:i+needleLength]) == needle {
					tmp = i
					break
				}
			} else {
				break
			}
		}
	}
	return tmp
}

// StrStr2 : the most despicable method
func StrStr2(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
