
def binarySearch(nums, target: int) -> int:
    lo = 0
    hi = len(nums)
    while lo < hi:
        print ('called')
        mid = (hi+lo)//2
        if target > nums[mid]:
            lo = mid + 1
        elif target == nums[mid]:
            return mid
        else:
            hi = mid
    if 0 <= lo < len(nums) and nums[lo] == target:
        return lo
    return -1
        
# print (search([-1,2,3,4,5,6,12], 4))