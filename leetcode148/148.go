/*
	leetcode tag:148 title:排序链表
*/

package leetcode148

type ListNode struct {
	Val  int
	Next *ListNode
}

// 抄的，不是自己做的
func SortList(head *ListNode) *ListNode {
	// 过滤条件
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head // 快慢指针
	var preSlow *ListNode    // 保存slow的前一个结点
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	preSlow.Next = nil  // 将中间节点的Next指向断开
	l := SortList(head) // 已排序的左链
	r := SortList(slow) // 已排序的右链

	return MergeList(l, r)
}

func MergeList(l1, l2 *ListNode) *ListNode {
	virtualNode := new(ListNode)
	tmp := virtualNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}

	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return virtualNode.Next
}
