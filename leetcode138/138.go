/*
	leetcode tag:138 title:复制带随机指针的链表
*/

package leetcode138

// 抄的，但是目前可以理解，用map来对应旧的Node和新的Node
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	// 条件过滤
	if head == nil {
		return nil
	}

	result := new(Node)
	tmp := result
	tmpHead := head
	m := map[*Node]*Node{} // key存储原head，value存储与原head的random对应的result

	for tmpHead != nil {
		tmp.Next = &Node{
			Val: tmpHead.Val,
		}
		m[tmpHead] = tmp.Next
		tmpHead = tmpHead.Next
		tmp = tmp.Next
	}

	tmp = result.Next
	tmpHead = head

	for tmpHead != nil {
		tmp.Random = m[tmpHead.Random]
		tmp = tmp.Next
		tmpHead = tmpHead.Next
	}

	return result.Next
}
