"""
- square the value and insert to position on array
"""
def sortedSquares(nums):
    left = 0
    right = len(nums) - 1
    res = []
    for i in range(len(nums)):
        if abs(nums[left]) >= abs(nums[right]):
            res.insert(0, nums[left] ** 2)
            left = left + 1
        else:
            res.insert(0, nums[right] ** 2)
            right = right - 1
    return res

print (sortedSquares([-4,-1,0,3,10]))