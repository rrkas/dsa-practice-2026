// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     carry := 0
//     var root, ans *ListNode
//     for l1 != nil || l2 != nil || carry > 0 {
//         n := &ListNode{}
//         s := carry
//         if l1 != nil {
//             s += l1.Val
//             l1 = l1.Next
//         }
//         if l2 != nil {
//             s += l2.Val
//             l2 = l2.Next
//         }
//         carry = s / 10
//         s = s % 10
//         n.Val = s
//         if root == nil {
//             root = n
//             ans = n
//         } else {
//             ans.Next = n
//             ans = n
//         }
//     }
//     return root
// }


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
        carry = sum / 10
    }

    return dummy.Next
}