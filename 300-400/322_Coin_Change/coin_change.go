package coinChainge

func coinChange(coins []int, amount int) int {
    if amount == 0 {
		return 0
	}

	needCoins := make([]int, amount+1)
	needCoins[0] = 0
	for i := 1; i < len(needCoins); i++ {
		min := amount+1
		for _, coin := range coins {
			if (i - coin) < 0 {
				continue
			}

			beforeCheck := i - coin
			if (beforeCheck == 0) {
				min = 0
				break
			}

			beforeCoins := needCoins[i-coin]
			if beforeCoins == 0 {
				continue
			}
			
			if (beforeCoins < min) {
				min = beforeCoins
			}
		}

		if (min != amount+1) {
			needCoins[i] = min + 1
		}
	}
	
	if needCoins[amount] == 0 {
		return - 1
	}
	return needCoins[amount]
}