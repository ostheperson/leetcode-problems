# class Node:
#     def __init__(self, val=None, nxt=None):
#         self.val = val
#         self.next = nxt

class Solution():
	def detectCycle(self, head: Node):
		slow, fast = head, head
		while fast and fast.next:
			fast = fast.next.next
			slow = slow.next
			if fast == slow:
				break
		if not fast or not fast.next:
			return None
		slow = head
		while slow != fast:
			slow = slow.next
			fast = fast.next
		return fast

