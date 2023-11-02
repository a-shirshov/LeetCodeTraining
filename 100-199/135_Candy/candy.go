package candy

func candy(ratings []int) int {
    candies := make([]int, 0, len(ratings))

	candies = append(candies, 1)
	for i := 1; i < len(ratings); i++ {
		candies = append(candies, 1)
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} 
	}

	for i := len(ratings) - 2; i >=0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] < candies[i+1] + 1 {
				candies[i] = candies[i+1] + 1
			}
		} 
	}

	var sum int
	for _, amountOfCandies := range candies {
		sum += amountOfCandies
	}
	return sum
}