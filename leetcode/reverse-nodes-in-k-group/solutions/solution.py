# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        prev = dummy

        while True:
            kth = prev
            for _ in range(k):
                kth = kth.next
                if not kth:
                    return dummy.next

            groupNext = kth.next

            prevNode = groupNext
            curr = prev.next

            while curr != groupNext:
                temp = curr.next
                curr.next = prevNode
                prevNode = curr
                curr = temp

            temp = prev.next
            prev.next = kth
            prev = temp
