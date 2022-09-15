def function_to_test():
	result = 0
	for i in range(0, 3):
		result += i
	return result

import time
# code to time how fast your algorithm runs
# get the start time
st = time.time()

function_to_test()

# get the end time
et = time.time()

# get the execution time
print('Execution time:', et - st, 'seconds')