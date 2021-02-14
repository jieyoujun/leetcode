package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min, profit := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if profit < prices[i]-min {
			profit = prices[i] - min
		}
	}
	return profit
	// 超时了baby
	// res, n := 0, len(prices)
	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		profit := (prices[j] - prices[i])
	// 		if res < profit {
	// 			res = profit
	// 		}
	// 	}
	// }
	// return res
}
