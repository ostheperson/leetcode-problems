class Solution:
    """
    dfs traversal would take in a tree
    intialte stack (toCheck) and a list (checked)
    add first item to toCheck stack
    while toCheck is not empty:
        pop toCheck stack
        check if popped value is in checked list
        if it is not then add to seen
        then
        for each adjacent node:
            if node is not in seen list:
                add to stack
    """
    def dfs(self, tree):
        toCheck = []
        checked = []
        toCheck.append(list(tree.keys())[0])
        while len(toCheck) > 0:
            nodeX = toCheck.pop()
            print (nodeX)
            if nodeX not in checked:
                checked.append(nodeX)
            for i in tree[nodeX]:
                if i not in checked:
                    toCheck.append(i)
        return checked

