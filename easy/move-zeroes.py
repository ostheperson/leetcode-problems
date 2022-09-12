def moveZeroes(nums):
    """
    Do not return anything, modify nums in-place instead.
    
    [1,2,8,9,0,0,9]
    [0,0,0,1,2,8,9]
    """
    zero = 0

    for index in range(1, len(nums)):
        if nums[zero] == 0:
            if nums[index] != 0:
                nums[zero], nums[index] = nums[index], nums[zero]
                zero += 1
        else:
            zero += 1
    return nums

print(moveZeroes([0,0,0,1,2,8,9]))