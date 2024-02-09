package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func([]int, int, int)
	dfs = func(combination []int, index int, total int) {
		if total == target {
			res = append(res, combination)
			return
		}

		if index >= len(candidates) || total > target {
			return
		}

		dfs(append(combination, candidates[index]), index, total+candidates[index])
		dfs(combination, index+1, total)
	}
	dfs([]int{}, 0, 0)
	return res
}
