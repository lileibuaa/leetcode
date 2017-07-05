package chapter21

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var startNode, goNode *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			if startNode == nil {
				startNode = l1
				goNode = l1
			} else {
				goNode.Next = l1
				goNode = goNode.Next
			}
			l1 = l1.Next
		} else {
			if startNode == nil {
				startNode = l2
				goNode = l2
			} else {
				goNode.Next = l2
				goNode = goNode.Next
			}
			l2 = l2.Next
		}
	}
	if l1 != nil {
		if startNode == nil {
			startNode = l1
		} else {
			goNode.Next = l1
		}
	} else if l2 != nil {
		if startNode == nil {
			startNode = l2
		} else {
			goNode.Next = l2
		}
	}
	return startNode
}
