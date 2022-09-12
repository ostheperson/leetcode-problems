"""
Input: head = [1,2,3,4,5,6]
            [1]
            [2,3]
            [3,5]
            [4,None]
Output: [3,4,5]
"""

def middleNode(head):
    slow, fast = head, head

    while (slow and fast.next):
        slow = slow.next
        fast = fast.next.next
    return slow
