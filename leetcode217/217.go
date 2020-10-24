/*
	leetcode tag:217 title:存在重复元素
*/

package leetcode217

import "sort"

// ContainsDuplicate : menthod of sort
func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := range nums {
		if i == len(nums)-1 {
			continue
		}
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// ContainsDuplicate2 : method of hashMap
func ContainsDuplicate2(nums []int) bool {
	hashMap := map[int]bool{}
	for _, num := range nums {
		if _, ok := hashMap[num]; ok {
			return true
		}
		hashMap[num] = true
	}
	return false
}
