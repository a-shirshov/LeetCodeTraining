package main

import "fmt"

func numToPos(num int, n int) [2]int {
	ost := num % n
	row := num / n
	if ost != 0 {
		row++
	}

	if ost == 0 {
		ost = n
	}
	col := 0

	if (row % 2 == 0) {
		col = n - ost
	} else {
		col = ost - 1
	}

	row = n - row

	return [2]int{row, col}
}

func snakesAndLadders(board [][]int) int {
	rows := len(board)
	cols := len(board[0])
	visited := make(map[int]struct{})
	startNum := 1
	queue := []int{}
	queue = append(queue, startNum)
	visited[startNum] = struct{}{}
	numberOfMoves := 0

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			elem := queue[0]
			queue = queue[1:]
			for j := 1; j < 7; j++ {
				newSquare := elem + j
				newPos := numToPos(newSquare, rows)

				if board[newPos[0]][newPos[1]] != -1 {
					newSquare = board[newPos[0]][newPos[1]]
				}
				
				if newSquare == rows*cols {
					return numberOfMoves + 1
				}

				_, hasVisit := visited[newSquare]
				if hasVisit {
					continue
				}

				queue = append(queue, newSquare)
				visited[newSquare] = struct{}{}
			}	
		}
		numberOfMoves++
	}
	return -1
}

func main() {
	fmt.Println(snakesAndLadders([][]int{
		{-1,-1,-1,-1,-1,-1},
		{-1,-1,-1,-1,-1,-1},
		{-1,-1,-1,-1,-1,-1},
		{-1,35,-1,-1,13,-1},
		{-1,-1,-1,-1,-1,-1},
		{-1,15,-1,-1,-1,-1},
	}))
}