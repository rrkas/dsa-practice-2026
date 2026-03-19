/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	before_head := &ListNode{}
	after_head := &ListNode{}

	before, after := before_head, after_head

	curr := head
	for curr != nil {
		if curr.Val < x {
			before.Next = curr
			before = curr
		} else {
			after.Next = curr
			after = curr
		}
		curr = curr.Next
	}

	after.Next = nil
	before.Next = after_head.Next
	return before_head.Next
}