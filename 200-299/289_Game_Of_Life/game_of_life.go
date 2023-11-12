package gameoflife

/*
	old | current | value
	0         0       0
	1         0       1
	0         1       2
	1         1       3
*/



func gameOfLife(board [][]int)  {
    rows, cols := len(board), len(board[0])

	countNeighbors := func(r, c int) int {
		nei := 0
		for i := r - 1; i < r + 2; i++ {
			for j := c - 1; j < c + 2; j++ {
				switch {
				case i==r && j==c:
					continue
				case i < 0 || j < 0:
					continue
				case i >= rows || j >= cols:
					continue
				case board[i][j] == 1 || board[i][j] == 3:
					nei++
				}
			}
		}
		return nei
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			nei := countNeighbors(i, j)

			if board[i][j] == 1 {
				if nei == 2 || nei == 3 {
					board[i][j] = 3
				}
			} else if nei == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 1 {
				board[i][j] = 0
			} else if board[i][j] == 2 || board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}