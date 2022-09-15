def twoSum(nums, target: int):
    l, r = 0, len(nums) - 1
    while(l != r):
        if (nums[l] + nums[r] > target):
            r -= 1
        elif (nums[l] + nums[r] < target):
            l += 1
        else:
            return [l+1, r+1]

print (twoSum([1,2,3,4,4,9,56,90], 8))