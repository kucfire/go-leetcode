/*
	leetcode tag:019 title:删除链表的倒数第N个结点
*/

package leetcode019

// ListNode : body
type ListNode struct {
	Val  int
	Next *ListNode
}

// normal method
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	second := &ListNode{0, head}
	for n > 0 {
		first = first.Next
		n--
	}
	result := second
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return result.Next
}

// method of double points
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	pre := head
	cur := head
	count := 0
	for pre != nil {
		count++
		if count > n {
			cur = cur.Next
		}
		pre = pre.Next
	}

	if count == n {
		return head
	}
	cur.Next = cur.Next.Next

	return result.Next
}
