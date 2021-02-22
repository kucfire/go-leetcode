package leetcode206

func ReverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		head, head.Next, res = head.Next, res, head
	}

	return res
}
