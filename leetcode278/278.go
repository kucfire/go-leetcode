/*
	leetcode tag:278 title:第一个错误的版本
*/

package leetcode278

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	head := 0
	tail := n
	mid := 0
	for head < tail {
		mid = head + (tail-head)/2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				return mid
			}
			tail = mid
		} else {
			head = mid + 1
		}
	}
	return 0
}
