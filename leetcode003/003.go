/*
	leetcode tag:003 title:无重复最长子字符串
*/

package leetcode003

//method of hashMap
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start, end, maxlength := 0, 0, 0
	array := make([]int, 0xffff)

	//reset
	for i := 0; i < len(array); i++ {
		array[i] = -1
	}

	a := []rune(s)

	//
	for ; end < len(a); end++ {
		value := array[a[end]]
		if value != -1 && value >= start {
			start = value + 1
		}

		array[a[end]] = end
		tmp := end - start
		if tmp >= maxlength {
			maxlength = tmp
		}
	}
	return maxlength + 1
}
