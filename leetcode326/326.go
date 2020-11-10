/*
	leetcode tag:326 title:3的幂
*/

package leetcode326

// IsPowerOfThree : normal method
func IsPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 2 {
		return false
	}
	for n > 3 {
		if n%3 != 0 {
			return false
		}

		n = n / 3
		if n < 3 {
			return false
		}
	}
	return true
}
