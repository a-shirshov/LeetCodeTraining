package main

type coord struct {
	r int
	c int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	visited := make(map[coord]struct{})
	rows := len(grid)
	cols := len(grid[0])
	islands := 0

	bfs := func(cd coord) {
		queue := []coord{}
		visited[cd] = struct{}{}
		queue = append(queue, cd)
		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for len(queue) != 0 {
			elem := queue[0]
			queue = queue[1:]
			for _, dr := range directions {	
				r := elem.r + dr[0]
				c := elem.c + dr[1]
				newCoord := coord{
					r: r,
					c: c,
				}
				_, visit := visited[newCoord]
				if (r < rows && r >= 0) && (c < cols && c >= 0) && grid[elem.r][elem.c] == '1' && !visit {
					queue = append(queue, newCoord)
					visited[newCoord] = struct{}{}
				}
			}
		}
		//fmt.Println(visited)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cd := coord{i, j}
			_, visit := visited[cd]
			if grid[i][j] == '1' && !visit {
				bfs(cd)
				islands++
			}
		}
	}
	return islands
}