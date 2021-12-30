#  Definition for singly-linked list.
class Node:
    def __init__(self, val=None, nxt=None):
        self.val = val
        self.next = nxt

class MyLinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    def addAtHead(self, val: int) -> None:
        """
        change self.head to new node and new node.next to former head
        """
        node = Node(val)

        node.next = self.head
        self.head = node

        self.size += 1
        return

    def get(self, index: int) -> int:
        if 0 <= index < self.size:
            curr = self.head
            for i in range(index):
                curr = curr.next
            return curr.val
        return -1

    def getNode(self, index: int) -> int:
        if 0 <= index < self.size:
            curr = self.head
            for i in range(index):
                curr = curr.next
            return curr
        return -1
        
        
    def addAtTail(self, val: int, nxt=None) -> None:
        if nxt and nxt < self.size:
            self.addAtIndex(self.size, val, nxt)
        else: self.addAtIndex(self.size, val)

    def addAtIndex(self, index: int, val: int, nxt=None) -> None:
        """
        loop through list till node before target index,
        point node before target index to new node,
        point new node to target
        """
        if 0 <= index <= self.size:
            if index == 0:
                self.addAtHead(val)
            else:
                curr = self.head
                pos = None
                for i in range(index-1):
                    if nxt and i == nxt:
                        pos = curr
                    curr = curr.next
                if nxt: node = Node(val, pos)
                else:
                    node = Node(val, curr.next)
                curr.next = node
                self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        """
        loop through till node before target index,
        point node before target index to target.next
        """
        if 0 <= index < self.size:
            curr = self.head
            if index == 0:
                self.head = curr.next
            else:
                curr = self.head
                for i in range(index-1):
                    curr = curr.next
                curr.next = curr.next.next
            self.size -= 1

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
            slow = slow.next
            if fast.next: fast = fast.next.next
            if slow == fast:
                return fast
        return False

    def detectCycle(self, head: Node) -> Node:
        count = 0
        d = head
        k = self.hasCycle(head)
        # print (k.val)
        while (d == k) == False:
            print ('iter: ' + str(count) + " d:" + str(d.val) + " k:" + str(k.val)  )
            d = d.next
            k = k.next
            count += 1
        return count


myList = MyLinkedList()
myList.addAtHead(0)
myList.addAtIndex(1, 1)
myList.addAtIndex(2, 2)
myList.addAtIndex(3, 3)
myList.addAtIndex(4, 4)
myList.addAtTail(5, 1)
sol = Solution()

# print (myList.getNode(0).val)
print (sol.hasCycle(myList.getNode(0)).val)
# print (sol.detectCycle(myList.getNode(0)))
