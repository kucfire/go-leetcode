/*
	leetcode tag:027 title:移除元素
*/
package leetcode027

func RemoveElement(nums []int, val int) int {
	mvIndex := 0
	for i, num := range nums {
		if num != val {
			nums[mvIndex] = nums[i]
			mvIndex++
		}
	}
	return mvIndex
}
