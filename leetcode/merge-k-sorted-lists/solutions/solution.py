# # Definition for singly-linked list.
# # class ListNode:
# #     def __init__(self, val=0, next=None):
# #         self.val = val
# #         self.next = next
# class Solution:
#     def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
#         dummy = ListNode()
#         root = dummy

#         while True:
#             minNode = None
#             nodeIdx = -1

#             for lidx, l in enumerate(lists):
#                 if l:
#                     if not minNode or l.val < minNode.val:
#                         minNode = l
#                         nodeIdx = lidx

#             if not minNode:
#                 break

#             dummy.next = minNode
#             dummy = dummy.next
#             lists[nodeIdx] = lists[nodeIdx].next

#         return root.next

###############################################################################

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
import heapq


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap = []

        for i, node in enumerate(lists):
            if node:
                heapq.heappush(heap, (node.val, i, node))

        dummy = ListNode()
        tail = dummy

        while heap:
            _, i, node = heapq.heappop(heap)

            tail.next = node
            tail = tail.next

            if node.next:
                heapq.heappush(heap, (node.next.val, i, node.next))

        return dummy.next
