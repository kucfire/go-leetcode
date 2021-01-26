/*
	leetcode tag:169 title:多数元素
	point:
	1.出现次数 大于 ⌊ n/2 ⌋ 的元素
	2.时间复杂度为 O(n)、空间复杂度为 O(1)
*/

package leetcode169

import (
	"sort"
)

// MajorityElementCount :method of count
func MajorityElementCount(nums []int) int {
	count, major := 0, 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if num == major {
			count++
		} else {
			count--
		}
	}
	return major
}

// MajorityElementHashmap :method of hashmap
func MajorityElementHashmap(nums []int) int {
	arrayMap := make(map[int]int)
	for _, num := range nums {
		arrayMap[num]++
	}
	for key, value := range arrayMap {
		if value > len(nums)/2 {
			return key
		}
	}
	return -1
}

// MajorityElementSort :排序
func MajorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// MajorityElementBit :位运算(不能用)
func MajorityElementBit(nums []int) int {
	result := 0
	signals := 0
	for i := 0; i < 32; i++ {
		ones := 0
		for _, num := range nums {
			if num < 0 {
				signals = 1
			}
			ones += (num >> i) & 1
			if ones > len(nums)/2 {
				result += 1 << i
				break
			}
		}
	}
	if result > 0 && signals == 1 {
		result = -result
	}
	return result
}
