/*
	leetcode tag:020 title:有效括号
*/

package leetcode020

// IsValid : stack method
func IsValid(s string) bool {
	lib := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	list := []rune{}
	// Str := []byte(s)

	for _, str := range s {
		if len(list) == 0 {
			list = append(list, str)
			continue
		}
		if str == '(' || str == '[' || str == '{' {
			list = append(list, str)
		} else if lib[str] == list[len(list)-1] {
			list = list[:len(list)-1]
		} else {
			list = append(list, str)
		}

	}

	if len(list) != 0 {
		return false
	}

	return true
}
