#  Definition for singly-linked list.
class Node:
    def __init__(self, val=None, nxt=None):
        self.val = val
        self.next = nxt

class Solution():
    def hasCycle(self, head: Node) -> bool:
        """ 
        initialize slow and fast pointers,
        loop if non of them are None
        if slow and fast.next are not the same, 
            set slow to next node and fast to next two nodes
        """
        slow = fast = head
        while slow and fast:
            if slow == fast.next:
                return True
            slow = slow.next
            if fast.next: fast = fast.next.next
        return False

# myList = MyLinkedList()
# myList.addAtHead(0)
# myList.addAtIndex(1, 1)
# myList.addAtIndex(2, 2)
# myList.addAtTail(3)

# sol = Solution()
# print (sol.hasCycle(myList.getNode(0)))