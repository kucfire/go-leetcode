/*
	leetcode tag:206 title:反转链表
*/

package leetcode206

type ListNode struct {
	Val  int
	Next *ListNode
}

// method of recursion
func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		head, head.Next, res = head.Next, res, head
	}
	return res
}
