/*
	leetcode tag:046 title:全排列
*/

package leetcode046

//normal method : recursion
func Permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	result := [][]int{}
	for i, num := range nums {
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[:i])
		copy(tmp[i:], nums[i+1:])

		tmp2 := Permute(tmp)
		for _, v := range tmp2 {
			result = append(result, append(v, num))
		}
	}
	return result
}
