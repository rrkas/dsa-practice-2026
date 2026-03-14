# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(
        self,
        list1: Optional[ListNode],
        list2: Optional[ListNode],
    ) -> Optional[ListNode]:
        if not list1 and list2:
            return list2

        if list1 and not list2:
            return list1

        if not list1 and not list2:
            return None

        dummy = ListNode()
        root = dummy
        if list1.val < list2.val:
            dummy.next = copy(list1)
            list1 = list1.next
        else:
            dummy.next = copy(list2)
            list2 = list2.next

        dummy = dummy.next

        while list1 and list2:
            if list1.val < list2.val:
                dummy.next = copy(list1)
                list1 = list1.next
                dummy = dummy.next
            else:
                dummy.next = copy(list2)
                list2 = list2.next
                dummy = dummy.next

        while list1:
            dummy.next = copy(list1)
            list1 = list1.next
            dummy = dummy.next

        while list2:
            dummy.next = copy(list2)
            list2 = list2.next
            dummy = dummy.next

        return root.next


def copy(node):
    return ListNode(val=node.val, next=node.next)
