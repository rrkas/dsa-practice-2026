/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	temp := head
	n := 0
	var last *ListNode
	for temp != nil {
		if temp.Next == nil {
			last = temp
		}
		temp = temp.Next
		n++
	}
	k %= n

	if k == 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head
	for i := 0; i < n-k; i++ {
		curr = curr.Next
		prev = prev.Next
	}
	dummy.Next = curr
	last.Next = head
	prev.Next = nil

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	tail := head

	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	k %= n
	if k == 0 {
		return head
	}

	// make circular
	tail.Next = head

	steps := n - k
	newTail := head

	for i := 1; i < steps; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}