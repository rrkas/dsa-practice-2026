/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		n1 := head
		n2 := head.Next

		prev.Next = n2
		n1.Next = n2.Next
		n2.Next = n1

		prev = n1
		head = n1.Next
	}

	return dummy.Next
}