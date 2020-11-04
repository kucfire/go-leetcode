/*
	leetcode tag:141 title:环形链表
*/

package leetcode141

type ListNode struct {
	Val  int
	Next *ListNode
}

// double points
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	// 从第一轮开始，需要以head的Next和head的Next.Next开始
	slow, fast := head.Next, head.Next.Next
	for {
		// 相等则退出
		if slow == fast {
			return true
		}
		// 非空则继续往前推，slow推一格，fast推两格
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
