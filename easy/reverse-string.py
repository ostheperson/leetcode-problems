def reverseString(s):
        """
        Do not return anything, modify s in-place instead.
        """
        l, r = 0,  len(s) - 1
        while (l < r):
            s[l], s[r] = s[r], s[l]
            l += 1
            r -= 1
        print (s)

print(reverseString(["h","e","l","l","o"]))