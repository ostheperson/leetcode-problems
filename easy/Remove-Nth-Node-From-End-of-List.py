class Solution:
    def removeNthFromEnd(self, head: Node, n: int):
        """
        set fast and slow to head
        shiht fast to by n positions
        increase slow and fast till fast.next = None
        set slow.next to node after nth node
        """
        fast = slow = head

        while (n > 0):
            n -= 1
            fast = fast.next
        if not fast:
            return head.next
        # print (fast.val)

        while fast.next != None:
            slow = slow.next
            fast = fast.next
        slow.next = slow.next.next
        return head