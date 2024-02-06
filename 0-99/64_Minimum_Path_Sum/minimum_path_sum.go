package minimumpathsum

func minPathSum(grid [][]int) int {
    var dp [][]int
	for i := 0; i < len(grid); i++ {
		dpline := []int{}
		for j := 0; j < len(grid[i]); j++ {
			if (i == 0) {
				if (j == 0) {
					dpline = append(dpline, grid[i][j])
				} else {
					dpline = append(dpline, grid[i][j] + dpline[j-1])
				}
				continue
			}

			if (j == 0) {
				dpline = append(dpline, grid[i][j] + dp[i-1][j])
				continue
			}
			
			min := dp[i-1][j]
			if (min > dpline[j-1]) {
				min = dpline[j-1]
			}
			min += grid[i][j]
			dpline = append(dpline, min)
		}
		dp = append(dp, dpline)
	}

	return dp[len(grid) - 1][len(grid[0]) - 1]
}