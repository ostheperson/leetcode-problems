# O(n+d)
def lengthOfLongestSubstring(s: str) -> int:
    """
    go right until you see char you have seen before,
    when you see char set left to one past char index.
    if char is not seen before, update max
    then put char in hash
    """

    if len(s) == 0: return 0
    if len(s) == 1: return 1
    left, right = 0, 0
    res = 0
    hash_book = {}
    while(right < len(s)):
        char = s[right]
        if char in hash_book and hash_book[char] >= left:
            left = hash_book[char] + 1
        else: 
            res = max(res, len(s[left:right])+1)
        hash_book[char] = right
        right += 1
    return res

print(lengthOfLongestSubstring("pwwkew"))
