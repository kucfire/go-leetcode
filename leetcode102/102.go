/*
	leetcode tag:102 title:二叉树的层数遍历
*/

package leetcode102

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// normal method
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	var loadTree func(node *TreeNode, i int)
	loadTree = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		// len必须-1才是与i同类型的存在，因为i是从0开始的，而len是真实长度
		// 这里是为了避免二元数组的第一层长度不够或者过长，以循环判断来进行
		for len(result)-1 < i {
			result = append(result, []int{})
		}
		result[i] = append(result[i], node.Val)
		loadTree(node.Left, i+1)
		loadTree(node.Right, i+1)
	}
	loadTree(root, 0)
	return result
}

//
