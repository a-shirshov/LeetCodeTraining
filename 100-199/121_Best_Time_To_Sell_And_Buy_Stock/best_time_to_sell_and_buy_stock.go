package besttimetosellandbuystock

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
    for i := 0; i < len(prices) - 1; i++ {
		if (prices[i] < minPrice) {
			minPrice = prices[i]
		}
		if (profit < prices[i+1]-minPrice) {
			profit = prices[i+1]-minPrice
		}
	}
	return profit
}