def binarySearch(nums, target: int, exclude) -> int:
    l = 0
    r = len(nums) - 1
    while l <= r:
        mid = (r + l) // 2
        if nums[mid] == target and mid != exclude:
            return mid
        elif nums[mid] > target:
            r = mid - 1
        else:
            l = mid + 1
    return -1

def twoSum(nums, target: int):
    # intiialize index to first item
    # get remainder from diff target and item at index
    # binary search nums for remainder
    # if value exist return the tuple array

    for i in range(len(nums)):
        diff = target - nums[i]
        ans = binarySearch(nums, diff, i)
        if (ans > -1):
            return [i+1, ans+1]

print (twoSum([1,2,3,4,4,9,56,90], 8))