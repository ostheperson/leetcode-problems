def getStartIndex(nums, target):
    lo, hi = 0, len(nums) - 1
    while (lo<=hi):
        mid = (lo + hi) >> 1
        if (nums[mid] > target):
            hi = mid - 1
        else:
            lo = mid + 1
    print (lo)
    return lo


def countOccurency(nums, target):
    # get start index of target
    # get send index of target+1
    startPosition = getStartIndex(nums, target)
    endPosition = getStartIndex(nums, target)
    return (endPosition-startPosition)

print (countOccurency([1,2,3,3,3,4], 3))
# print (countOccurency([1,2,3,3,3,4,5,6,7,7,7,7,8], 3))