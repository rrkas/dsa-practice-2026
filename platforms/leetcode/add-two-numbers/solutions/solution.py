# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        carry = 0
        root, ans = None, None
        while l1 != None or l2 != None or carry > 0:
            s = carry
            if l1 != None:
                s += l1.val
                l1 = l1.next
            if l2 != None:
                s += l2.val
                l2 = l2.next
            carry = s // 10
            s = s % 10
            n = ListNode(val=s)
            if root is None:
                root = n
                ans = n
            else:
                ans.next = n
                ans = n

        return root
