def searchInsert(nums, target: int) -> int:
    """
    - find first item bigger than target
    - return index
    """
    left = 0
    right = len(nums) - 1
    while (left <= right):
        mid = (left + right) >> 1
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return left
    

print (searchInsert([1,2,3], 4))