/*
	leetcode tag:204 title:计数质数
*/

package leetcode204

// normal method
func countPrimes(n int) int {
	identi := make([]bool, n+1)
	var count int
	for i := 2; i < n; i++ {
		if identi[i] {
			continue
		}
		for j := i; j < n; j += i {
			identi[j] = true
		}
		count++
	}
	return count
}
