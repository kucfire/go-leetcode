/*
	leetcode tag:287 title:寻找重复数
*/

package leetcode287

// method of myself
func FindDuplicate(nums []int) int {
	list := make([]int, len(nums))
	for _, num := range nums {
		list[num]++
		if list[num] > 1 {
			return num
		}
	}
	return -1
}

// 快慢指针
func FindDuplicate2(nums []int) int {
	slow, fast := 0, 0
	// 确定该数组内的元素可否形成一个闭环
	for slow, fast = nums[slow], nums[nums[fast]]; fast != slow; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
