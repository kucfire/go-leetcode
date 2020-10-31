/*
	leetcode tag:007 title:整数反转
*/

package leetcode007

// Reverse : normal method
func Reverse(x int) int {
	result := 0
	// tmp := 0
	for x != 0 {
		if temp := int32(result); (temp*10)/10 != temp {
			return 0
		}
		result = result*10 + x%10
		x = x / 10
	}
	return result
}
