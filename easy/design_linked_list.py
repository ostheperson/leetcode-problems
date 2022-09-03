class Node:
	def __init__(self, val=None, nxt=None):
		self.val = val
		self.next = nxt

class LinkedList:
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
		
	def getNode(self, index: int):
		if 0 <= index < self.size:
			curr = self.head
			for i in range(index):
				curr = curr.next
			return curr
		return None 
		
	def addAtTail(self, val: int) -> None:
		self.addAtIndex(self.size, val)

	def addAtIndex(self, index: int, val: int) -> None:
		"""
		loop through list till node before target,
		point node before target to new node,
		point new node to target
		"""
		if 0 <= index <= self.size:
			if index == 0:
				self.addAtHead(val)
			else:
				curr = self.head
				for i in range(index-1):
					curr = curr.next

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

