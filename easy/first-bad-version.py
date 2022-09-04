def isBadVersion(num):
    arr = [False, True, True, True, True]
    return arr[((num-1) % 5)]

def firstBadVersion(n: int) -> int:
    left = 1
    right = n
    while left < right:
        print (str(left) + ' ' + str(right))
        mid = (left + right) // 2
        if isBadVersion(mid):
            right = mid 
        else:
            left = mid + 1
    return right

        
print (firstBadVersion(5))