/*
	leetcode tag:101 title:对称二叉树
*/

package leetcode101

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// normal method
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var checkValue func(left, right *TreeNode) bool
	checkValue = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left != nil && right != nil {

			if left.Val == right.Val {
				return checkValue(left.Left, right.Right) == true && checkValue(left.Right, right.Left) == true
			}
		}
		return false
	}
	return checkValue(root.Left, root.Right)
}
