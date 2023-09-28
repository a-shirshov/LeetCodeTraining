package validsudoku

import "fmt"

func isValidSudoku(board [][]byte) bool {
	size := 9
	hashMap := make(map[string]bool)
    for i:=0; i<size; i++{
		for j:=0; j<size;j++{
			currentValue := string(board[i][j])
			if currentValue == "."{
				continue
			}
			
			_,ok1 := hashMap[currentValue + "in row" + fmt.Sprint(i)]
			_,ok2 := hashMap[currentValue + "in column" + fmt.Sprint(j)]
			_,ok3 := hashMap[currentValue + "in grid" + fmt.Sprint(i/3) + fmt.Sprint(j/3)]

			if ok1 || ok2 || ok3 {
				return false
			}

			hashMap[currentValue + "in row" + fmt.Sprint(i)] = true
			hashMap[currentValue + "in column" + fmt.Sprint(j)] = true
			hashMap[currentValue + "in grid" + fmt.Sprint(i/3) + fmt.Sprint(j/3)] = true
		}
	}
	return true
}