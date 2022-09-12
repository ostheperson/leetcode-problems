def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head.next:
            return None
        slow, fast = head, head
        prev = head
        while (fast != None and fast.next != None):
            prev = slow
            slow = slow.next
            fast = fast.next.next
            
        prev.next = slow.next
        slow.next = None
        return head