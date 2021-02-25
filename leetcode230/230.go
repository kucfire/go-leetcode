/*
	leetcode tag:230 title:二叉搜索树中第K小的元素
*/

package leetcode230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 抄的
func KthSmallest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(root *TreeNode, k *int)
	dfs = func(root *TreeNode, k *int) {
		if root != nil {
			dfs(root.Left, k)
			*k--
			if *k == 0 {
				ans = root.Val
			}
			dfs(root.Right, k)
		}
	}
	dfs(root, &k)
	return ans
}

// 非递归
func KthSmallest2(root *TreeNode, k int) int {
	res := 0
	stack := []*TreeNode{}
	var cur *TreeNode = root

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}

	return res
}
