# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if head is None or head.next is None or k == 0:
            return head

        n = 1
        tail = head
        while tail.next is not None:
            tail = tail.next
            n += 1

        k %= n
        if k == 0:
            return head

        tail.next = head
        newtail = tail
        for i in range(n - k):
            newtail = newtail.next

        newhead = newtail.next
        newtail.next = None

        return newhead
