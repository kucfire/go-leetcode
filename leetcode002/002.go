/*
	leetcode tag:002 title:两数之和
*/

package leetcode002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	need to solve problem
	1. zero value
	2. into a digit(进位取值)
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := new(ListNode)
	result := tmp
	up := 0
	for l1 != nil || l2 != nil || up != 0 {
		// (if l1 != nil) mean l1 have a Val to calculation and l1.Val will add in up
		if l1 != nil {
			up += l1.Val
			l1 = l1.Next
		}
		// (if l2 != nil) mean l2 have a val to calculation and l2.Val will add in up
		if l2 != nil {
			up += l2.Val
			l2 = l2.Next
		}
		// 取余数
		tmp.Next = &ListNode{up % 10, nil}
		up = up / 10
		tmp = tmp.Next
	}
	return result.Next
}
