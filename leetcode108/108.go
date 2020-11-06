/*
	leetcode tag:108 title:将有序数组转换成二叉搜索树
*/

package leetcode108

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// normal method
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]
	node := &TreeNode{nums[mid], sortedArrayToBST(left), sortedArrayToBST(right)}
	return node
}
