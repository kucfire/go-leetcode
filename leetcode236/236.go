/*
	leetcode tag:236 title:二叉树的最近公共祖先
*/

package leetcode236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 抄的
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// var pAncestor, qAncestor []*TreeNode

	if root == nil {
		return nil
	}

	if root.Val == q.Val || root.Val == p.Val {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left

}
