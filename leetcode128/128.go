/*
	leetcode tag:128 title:最长连续序列
*/

package leetcode128

func LongestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	total := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if total < currentStreak {
				total = currentStreak
			}
		}
	}
	return total
}
