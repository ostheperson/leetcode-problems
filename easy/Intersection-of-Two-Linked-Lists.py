# class Node:
#     def __init__(self, val=None, nxt=None):
#         self.val = val
#         self.next = nxt
   
class Solution:
    # O(1) memory 0(n + m) time
    def getIntersectionNode(self, headA: Node, headB: Node) -> Node:
        i = headA; j = headB
        
        while i != j:
            i = i.next if i else headB
            j = j.next if j else headA
        return i