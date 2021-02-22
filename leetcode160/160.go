/*
	leetcode tag:160 title:相交链表
*/

package leetcode160

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := 0
	tmpA := headA
	for tmpA != nil {
		lengthA++
		tmpA = tmpA.Next
	}

	lengthB := 0
	tmpB := headB
	for tmpB != nil {
		lengthB++
		tmpB = tmpB.Next
	}

	if lengthA > lengthB {
		for lengthA != lengthB {
			headA = headA.Next
			lengthA--
		}
	} else {
		for lengthB != lengthA {
			headB = headB.Next
			lengthB--
		}
	}

	// 用headA或者headB做判断
	for headA != nil {
		if headA == headB {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}

	return nil
}
