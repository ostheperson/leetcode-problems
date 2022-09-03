class Solution:
    """
    tree = { 
            'a': {'b', 'c'},
            'b': {'a'}
            'c': {'a'}
    }
    """
    def bfs(self, tree):
        toCheck = []
        checked = []
        toCheck.append(list(sorted(tree.keys()))[0])
        while len(toCheck) > 0:
            nodeX = toCheck.pop(0)
            if nodeX not in checked:
                checked.append(nodeX)
            for i in tree[nodeX]:
                if i not in checked:
                    toCheck.append(i)
        return checked

sol = Solution()
tree = {'1': ['3'], '3': ['1', '2', '4'], '2': ['3'], '4': ['3', '5'], '5': ['4', '6'], '6': ['5', '7', '8'], '7': ['6'], '8': ['6']}
print (sol.bfs(tree))