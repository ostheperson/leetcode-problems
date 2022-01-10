# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    # better memory
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        seen = []
        check = []
        
        check.append(root)
        
        while len(check) > 0:
            node = check.pop()
            if node:
                if node.right not in check:
                    check.append(node.right)
                if node.left != None and node.left not in seen:
                    check.append(node)
                    check.append(node.left)
                elif node.left == None or node.left in seen:
                    seen.append(node)
                
        return [x.val for x in seen]

    # better speed
    def inorderTraversalRecursive(self, root: Optional[TreeNode]) -> List[int]:
        result = []
        
        def inorder(node):
            if node:
                inorder(node.left)
                result.append(node.val)
                inorder(node.right)

        inorder(root)
        return result