class Solution:
    #       1->2->3->None
    # None<-1<-2<-3
    def reverseLinkedListRecursive(self, head, prev=None) -> Node:
        t = head.next
        head.next = prev
        prev = head
        if (t == None):
            return head
        return self.reverseLinkedListRecursive(t, prev)

    def reverseLinkedList(self, head, prev=None) -> Node:
        prev = None
        curr = head
        while (curr):
            t = curr.next      
            curr.next = prev
            prev = curr
            curr = t
        return prev
