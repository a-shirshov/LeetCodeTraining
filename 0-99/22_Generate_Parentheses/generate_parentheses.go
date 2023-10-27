package generateparentheses

func backtracking (resultArray *[]string, currentString string, openCount, closedCount, n int) {
	if len(currentString) == n*2 {
		*resultArray = append(*resultArray, currentString)
	} 
	if openCount < n {
		backtracking(resultArray, currentString + "(", openCount + 1, closedCount, n)
	}
	if closedCount < openCount {
		backtracking(resultArray, currentString + ")", openCount, closedCount + 1, n)
	}
}

func generateParenthesis(n int) []string {
    var result []string

	backtracking(&result, "", 0, 0, n)

	return result
}