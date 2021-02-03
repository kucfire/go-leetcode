package leetcode242

import "reflect"

func IsAnagramReview(s string, t string) bool {
	sMap := make(map[rune]int)
	for _, str := range s {
		sMap[str]++
	}
	tMap := make(map[rune]int)
	for _, str := range t {
		tMap[str]++
	}
	if !reflect.DeepEqual(sMap, tMap) {
		return false
	}
	return true
}
