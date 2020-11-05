/*
	leetcode tag:104 title:二叉树的最大深度
*/

package leetcode104

// TreeNode : body
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归 resursion
func maxDepth(root *TreeNode) int {
	var depth func(*TreeNode, int) int
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	depth = func(node *TreeNode, count int) int {
		if node == nil {
			return count
		}
		count = max(depth(node.Left, count+1), depth(node.Right, count+1))
		return count
	}
	result := depth(root, 0)
	return result
}
