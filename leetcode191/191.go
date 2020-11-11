/*
	leetcode tag:191 title:位一的个数
*/

package leetcode191

// HammingWeight :
func HammingWeight(num uint32) int {
	var acount int

	for ; num > 0; acount++ {
		num = num & (num - 1)
	}

	return acount
}

// HammingWeight2 :
func HammingWeight2(num uint32) int {
	var acount int
	for num > 0 {
		if num%2 == 1 {
			acount++
		}
		num = num / 2
	}
	return acount
}
