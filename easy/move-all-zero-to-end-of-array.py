
def moveZero(nums):
	count = 0
	for i in range(len(nums)):
		if nums[i] != 0:
			nums[count] = nums[i]
			count += 1
	for i in range(count, len(nums)):
		nums[i] = 0
	return nums

print(moveZero([1, 2, 0, 0, 0, 3, 6]))