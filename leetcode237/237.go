/*
	leetcode tag:237 title:删除链表中的节点
*/

package leetcode237

// ListNode Nody
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteNode : normal method
func DeleteNode(node *ListNode) {
	*node = *node.Next
}
