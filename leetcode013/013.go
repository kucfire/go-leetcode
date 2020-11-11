/*
	leetcode tag:013 title:罗马数字转整数
*/

package leetcode013

// RomanToInt : normal method
func RomanToInt(s string) int {
	luoma := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// luoma := make([]int, 0xffff)
	// luoma['I'] = 1
	// luoma['V'] = 5
	// luoma['X'] = 10
	// luoma['L'] = 50
	// luoma['C'] = 100
	// luoma['D'] = 500
	// luoma['M'] = 1000

	var result int
	// signed := false
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			result += luoma[s[i]]
		} else {
			if luoma[s[i]] >= luoma[s[i+1]] {
				result += luoma[s[i]]
			} else {
				result -= luoma[s[i]]
			}
		}
	}
	return result
}
