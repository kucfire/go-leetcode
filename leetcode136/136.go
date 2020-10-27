/*
	leetcode tag:136 title:只出现一次的数字
*/

package leetcode136

//
func SingleNumber(nums []int) int {
	arrayList := make(map[int]int)
	for _, num := range nums {
		arrayList[num]++
	}
	for i, num := range arrayList {
		if num == 1 {
			return i
		}
	}
	return 0
}

func SingleNumber2(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
