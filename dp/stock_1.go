package dp

func maxProfit(prices []int) int {

	start := 0
	max_profit := 0

	for i, _ := range prices {
		if prices[i] > prices[start] {
			if prices[i]-prices[start] > max_profit {
				max_profit = prices[i] - prices[start]
			}
		}
		if prices[i] < prices[start] {
			start = i
		}
	}
	return max_profit

}
