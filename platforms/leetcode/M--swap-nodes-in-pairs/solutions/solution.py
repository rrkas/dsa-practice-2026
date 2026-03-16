# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        prev = dummy

        while head and head.next:
            n1 = head
            n2 = head.next

            prev.next = n2
            n1.next = n2.next
            n2.next = n1

            prev = n1
            head = n1.next

        return dummy.next
