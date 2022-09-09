"""
- get index
- get shift right by k
- switch positon with item in shifted place
[1,1,2,3,4,5,6]

[7,1,2,3,4,5,6]

[6,7,1,2,3,4,5]
"""

def moveByOne(nums):
    last_item = nums[len(nums)-1]
    for i in range(len(nums) - 1, 0, -1):
        nums[i] = nums[i-1]
    nums[0] = last_item

def rotate(nums, k: int):
    for i in range(k):
        moveByOne(nums)

    print (nums)


print (rotate([1,2,3,4,5,6,7], 2))

