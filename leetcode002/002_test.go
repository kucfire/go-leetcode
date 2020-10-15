// package leetcode002

// import "testing"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func TestaddTwoNumbers(t *testing.T) {
// 	var l1_1 = new(ListNode)
// 	var l1_2 = new(ListNode)
// 	var l1_3 = new(ListNode)
// 	l1_1.Val = 2
// 	l1_2.Val = 4
// 	l1_3.Val = 3
// 	l1_1.Next = l1_2
// 	l1_2.Next = l1_3

// 	var l2_1 = new(ListNode)
// 	var l2_2 = new(ListNode)
// 	var l2_3 = new(ListNode)
// 	var l2_4 = new(ListNode)
// 	l2_1.Val = 5
// 	l2_2.Val = 6
// 	l2_3.Val = 4
// 	l2_4.Val = 2
// 	l2_1.Next = l2_2
// 	l2_2.Next = l2_3
// 	l2_3.Next = l2_4

// 	result := new(ListNode)
// 	result.Val = 7
// 	result.Next = new(ListNode)
// 	result.Next.Val = 0
// 	result.Next.Next = new(ListNode)
// 	result.Next.Next.Val = 8
// 	result.Next.Next.Next = new(ListNode)
// 	result.Next.Next.Next.Val = 2

// 	actual := addTwoNumbers(l1_1, l2_1)
// 	for actual != nil && result != nil {
// 		if actual.Val != result.Val {
// 			t.Errorf("value is error")
// 		}
// 	}
// }
