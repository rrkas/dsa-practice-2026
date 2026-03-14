/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		kth := prev
		for _ = range k {
			kth = kth.Next
			if kth == nil {
				return dummy.Next
			}
		}
		// end of loop -> kth contains last node of group

		groupNext := kth.Next
		prevNode := groupNext
		curr := prev.Next

		for curr != groupNext {
			temp := curr.Next
			curr.Next = prevNode
			prevNode = curr
			curr = temp
		}

		temp := prev.Next
		prev.Next = kth
		prev = temp
	}

	return dummy.Next
}