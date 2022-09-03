from design_linked_list import LinkedList
"""
0 -> 3 
0 -> 1

xx -> 0 -> 0 -> 
"""
class Node:
    def __init__(self, val="", nxt=None):
        self.val = val
        self.next = nxt

def mergeTwoLists(list1, list2):
    dummy = Node()

    tail = dummy

    while (True):
        if list1 is None:
            tail.next = list2
            break
        if list2 is None:
            tail.next = list1
            break

        if list1.val <= list2.val:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = tail.next

    return dummy.next


list1 = LinkedList()
list1.addAtTail(2)
list1.addAtTail(3)

list2 = LinkedList()
list2.addAtTail(1)
list1.addAtTail(1)

print(mergeTwoLists(list1.getNode(0), list2.getNode(0)).val)
        