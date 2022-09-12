def binarySearch(nums, target: int) -> int:
    """
    1. set left and right
    2. if left > right: terminate
    3. set mid
    4. if mid < targat: set left to mid + 1
    5. if mid > targat: set right to mid - 1
    6. if mid == target return
    """
    l = 0
    r = len(nums) - 1
    while l <= r:
        mid = (r + l) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] > target:
            r = mid - 1
        else:
            l = mid + 1
    return -1


print (binarySearch([2,7,11,15], 7))