/*
	leetcode tag:328 title:奇偶链表
*/

package leetcode328

type ListNode struct {
	Val  int
	Next *ListNode
}

// 题解答案
func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	evenTMP := even

	for evenTMP != nil && evenTMP.Next != nil {
		odd.Next = evenTMP.Next
		odd = odd.Next
		evenTMP.Next = odd.Next
		evenTMP = evenTMP.Next
	}
	odd.Next = even

	return head
}

// 不行
func OddEvenListMyself(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	var evenTMP *ListNode = nil

	for odd != nil && odd.Next != nil {
		evenTMP = odd.Next
		evenTMP = evenTMP.Next
		odd = evenTMP.Next
		odd = odd.Next
	}
	odd.Next = even

	return head
}
