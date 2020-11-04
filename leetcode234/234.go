/*
	leetcode tag:234 title:回文链表
*/

package leetcode234

type ListNode struct {
	Val  int
	Next *ListNode
}

// method of stack
func isPalindrome(head *ListNode) bool {
	arrayList := []*ListNode{}
	for head != nil {
		arrayList = append(arrayList, head)
		head = head.Next
	}

	for i, j := 0, len(arrayList)-1; i <= j; i, j = i+1, j-1 {
		if arrayList[i].Val != arrayList[j].Val {
			return false
		}
	}
	return true
}
