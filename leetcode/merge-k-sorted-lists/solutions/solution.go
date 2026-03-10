// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func mergeKLists(lists []*ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	root := dummy

// 	for {
// 		var minNode *ListNode = nil
// 		var lidx int
// 		for i, list := range lists {
// 			if list != nil {
// 				if minNode == nil || list.Val < minNode.Val {
// 					minNode = list
// 					lidx = i
// 				}
// 			}
// 		}

// 		if minNode == nil {
// 			break
// 		}

// 		dummy.Next = minNode
// 		lists[lidx] = lists[lidx].Next
// 		dummy = dummy.Next
// 	}

// 	return root.Next
// }


func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    for len(lists) > 1 {
        merged := []*ListNode{}

        for i := 0; i < len(lists); i += 2 {
            l1 := lists[i]
            var l2 *ListNode

            if i+1 < len(lists) {
                l2 = lists[i+1]
            }

            merged = append(merged, mergeTwo(l1, l2))
        }

        lists = merged
    }

    return lists[0]
}

func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }

    if l1 != nil {
        tail.Next = l1
    } else {
        tail.Next = l2
    }

    return dummy.Next
}