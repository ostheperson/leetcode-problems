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
        toCheck.append(list(tree.keys())[0])
        while len(toCheck) > 0:
            nodeX = toCheck.pop(0)
            if nodeX not in checked:
                checked.append(nodeX)
            for i in tree[nodeX]:
                if i not in checked:
                    toCheck.append(i)
        return checked

