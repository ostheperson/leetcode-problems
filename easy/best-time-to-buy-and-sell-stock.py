class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        """
        1. set initial price, min = prices [0] and profit = 0
        2. get curr item
            check if curr price is smaller than previous if true set min
        3. get difference with previous day's price
            curr price - min
                if larger than profit set profit
        4. repeat 2
        5. return profit
        """
        init = minPrice = prices[0]
        profit = 0
        for i in prices:
            if i < minPrice:
                minPrice = i
            if i - minPrice > profit:
                profit = i - minPrice
        return profit
        