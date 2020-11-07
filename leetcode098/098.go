/*
	leetcode tag:098 title:验证二叉搜索树
*/

package leetcode098

import "math"

// TreeNode : body
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// method of recursion(递归)
func isValidBST(root *TreeNode) bool {
	var decide func(*TreeNode, int, int) bool
	decide = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}

		return decide(root.Left, lower, root.Val) && decide(root.Right, root.Val, upper)
	}

	return decide(root, math.MinInt64, math.MaxInt64)
}

// method of stack by Middle order traversal
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		// 把当前root往下所有最左侧的node放进栈里面，取最后一个，即最小的node值
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
