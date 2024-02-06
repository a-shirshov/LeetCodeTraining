package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    var dp [][]int
	for i := 0; i < len(obstacleGrid); i++ {
		var dpline []int
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dpline = append(dpline, 0)
				continue
			}

			if (i == 0) {
				if (j == 0) {
					dpline = append(dpline, 1)
					continue
				}

				dpline = append(dpline, dpline[j-1])
				continue
			}

			if (j == 0) {
				dpline = append(dpline, dp[i-1][j])
				continue
			}

			dpline = append(dpline, dp[i-1][j]+dpline[j-1])
		}
		dp = append(dp, dpline)
	}
	return dp[len(dp) - 1][len(dp[0]) - 1]
}