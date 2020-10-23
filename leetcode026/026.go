/*
	leetcode tag:026 title:删除数组中的重复项
*/

package leetcode026

// RemoveDuplicates : method of normal
func RemoveDuplicates(nums []int) int {
	maxLength := 0
	if len(nums) == 0 {
		return maxLength
	}
	for i, num := range nums {
		if i == 0 {
			maxLength++
			continue
		}
		//寻找下一个非重复元素，并添加到maxlength对应的位置上
		if num != nums[i-1] {
			nums[maxLength] = num // 更改原数组的数值
			maxLength++
		} else {
			continue
		}
	}
	return maxLength
}
