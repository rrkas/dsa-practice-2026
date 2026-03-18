/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	gprev := dummy
	curr := head

	for curr != nil {
		// check if duplicate sequence starts
		if curr.Next != nil && curr.Val == curr.Next.Val {
			val := curr.Val

			// skip all duplicates
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}

			gprev.Next = curr
		} else {
			gprev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}