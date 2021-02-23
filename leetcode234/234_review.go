package leetcode234

func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast, slow := head, head
	// go的判断条件要分先后顺序 *
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	halfStart := slow.Next
	var reverse *ListNode = nil
	for halfStart != nil {
		halfStart, halfStart.Next, reverse = halfStart.Next, reverse, halfStart
	}

	p1 := head
	p2 := reverse
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}
