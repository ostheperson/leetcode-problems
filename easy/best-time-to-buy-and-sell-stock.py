def maxProfit(prices) -> int:
	""" 
	- max positive difference in price between days
	- find when to by 
	- find when to sell
	- identify lower stock points

	traverse each item
	check if it is the lowest
		store index
		continue
	get differences between lowest and current
	store index with max difference
	"""
	lowest_price = prices[0]
	max_difference = 0

	for price in prices:
		if price < lowest_price:
			lowest_price = price
		else:
			difference = price - lowest_price
			if (max_difference < difference):
				max_difference = difference

	return max_difference

print(maxProfit([7,3,2,4]))



		