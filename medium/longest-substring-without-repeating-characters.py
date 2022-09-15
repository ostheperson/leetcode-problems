def lengthOfLongestSubstring(s: str) -> int:
    stack = []
    max_lenght = 0
    for i in s:
        if i not in stack:
            stack.append(i)
        else: 
            stack = []
        max_lenght = max(max_lenght, len(stack))
    return max_lenght

print(lengthOfLongestSubstring('pwwkew'))