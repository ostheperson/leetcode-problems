"""
	Given an array of integers nums and an integer k, 
	return the total number of continuous subarrays 
	whose sum equals to k.
	
	num = [3,-3,5,0,1] k = 0
	output = 2
"""
def subarraySum(nums, k: int) -> int:
	

print (subarraySum([0, 3, -3, 0, 1, 2], 3))
	

"""
[3,1,2,3]
1. go through list
2. check for k
3. if we find increase count
4. add value to some sum
5. if sum == k increase count
6. return count
"""

"""
find multiples of 3 under 1000 
find multilpes of 5 under 1000
starting at 3 and adding all multiples after din

multiples = 0

for i < 1000:
	multiples += i if i % 3 == 0 else i if i%5 == 0:
print (multiple)
"""