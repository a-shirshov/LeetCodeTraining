package main

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	
	var result []int
	currentN := 0
	currentM := 0
	for i := 0; n > 0 && m > 0; i++ {
		switch i % 4 {
		case 0:
			for j := 0; j < m; j++ {
				result = append(result, matrix[currentN][currentM])
				currentM++
			}
			currentM--
			currentN++
			n--
		case 1:
			for j := 0; j < n; j++ {
				result = append(result, matrix[currentN][currentM])
				currentN++
			}
			currentN--
			currentM--
			m--
		case 2:
			for j := 0; j < m; j++ {
				result = append(result, matrix[currentN][currentM])
				currentM--
			}
			currentM++
			currentN--
			n--
		case 3:
			for j := 0; j < n; j++ {
				result = append(result, matrix[currentN][currentM])
				currentN--
			}
			currentN++
			currentM++
			m--
		}
	}
	return result
}

func main() {
	spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}})
}