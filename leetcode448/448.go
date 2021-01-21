/*
	leetcode tag:448 title:找到所有数组中消失的数字
*/

package leetcode448

// myself
func FindDisappearedNumbers(nums []int) []int {
	var result []int
	// 将遍历到的数标记在value
	numsMap := map[int]int{}
	for _, num := range nums {
		numsMap[num] = 1
	}

	//
	for j := 1; j <= len(nums); j++ {
		if numsMap[j] != 1 {
			result = append(result, j)
		}
	}
	return result
}

// FindDisappearedNumbers2 : new method
func FindDisappearedNumbers2(nums []int) []int {
	var result []int

	for _, num := range nums {
		// 遍历过程中，某些数会被提前修改成负值，需要将值拷贝出来的值改为绝对值，再进行下一步的逻辑判断
		if num < 0 {
			num = -num
		}
		// 修改num-1对应下标的值为负值
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	// 查找大于0的数
	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
