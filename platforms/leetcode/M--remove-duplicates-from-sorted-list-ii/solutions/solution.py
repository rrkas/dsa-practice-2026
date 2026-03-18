# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        gprev = dummy
        curr = head
        while curr is not None:
            if curr.next is not None and curr.val == curr.next.val:
                val = curr.val

                while curr is not None and curr.val == val:
                    curr = curr.next

                gprev.next = curr
            else:
                gprev = curr
                curr = curr.next

        return dummy.next
