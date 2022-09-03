def isValid(s: str) -> bool:
    types = {
        '{': '}', 
        '[': ']',
        '(': ')'
    }
    stack = []
    # s = '(()[{}])'
    for i in s:
        if i in types:
            stack.append(i)
        elif not stack or types[stack.pop()] != i:
            return False

    return len(stack) == 0

print (isValid("(()[{}])"))

