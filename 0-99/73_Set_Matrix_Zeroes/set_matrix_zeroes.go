package setmatrixzeroes

func setZeroes(matrix [][]int)  {
    hashMapCursedi := make(map[int]struct{})
	hashMapCursedj := make(map[int]struct{})

	cursej := func(i, j int) {
		for indexj := j; indexj >=0; indexj-- {
			matrix[i][indexj] = 0
		}
	}

	cursei := func(i, j int) {
		for indexi := i; indexi >=0; indexi-- {
			matrix[indexi][j] = 0
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			_, oki := hashMapCursedi[i] 
			_, okj := hashMapCursedj[j]
			if matrix[i][j] == 0 {
				cursei(i,j)
				cursej(i,j)
				hashMapCursedi[i] = struct{}{}
				hashMapCursedj[j] = struct{}{}
				continue
			}
			
			if oki || okj {
				matrix[i][j] = 0
			}
		}
	}
}