
def subarraySum(nums, k: int) -> int:
	count = 0
	sums = 0
	for i in range(len(nums)):
		sums += nums[i]
		if sums == k: count += 1
		elif nums[i] == k: count += 1
		
	return count

print (subarraySum([1, 1, 1], 2))



# def reverser(arr):
#     for i in range(len(arr)):
#         temp = arr[i]
#         arr[i] = arr[(len(arr)-1-i)]
#         print
#         arr[(len(arr)-1-i)] = temp
#     return arr

# def invertImage(image):
#     # Write your code here
#     for row in range(len(image)):
#         image[row] = reverser(image[row])
#     return reverser(image)
#     # return ' '.join(sentence)

# print (invertImage([[7,2,3], [7,7,1]]))