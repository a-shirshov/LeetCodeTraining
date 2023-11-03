package zigzagconversion

import "strings"

//PAYPALISHIRING
// 0 2 4 6 8 10 12 -> 1 3 5 7 9 11 13 (2)
// 0 4 8 12 -> 1 3 5 7 9 11 13 -> 2 6 10 (3)
// 0 6 12 -> 1 5 7 11 13 -> 2 4 8 10 -> 3 9 (4)


func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	maxStep := (numRows - 1) * 2

	var result strings.Builder
    for i := 0; i < numRows; i++ {
		leftLeg := maxStep - i * 2
		rightLeg := maxStep - leftLeg
		if leftLeg == 0 {
			leftLeg = rightLeg
		}

		if rightLeg == 0 {
			rightLeg = leftLeg
		}
		count := 0
		index := i
		for index < len(s) {
			result.WriteByte(s[index])
			if count % 2 == 0 {
				index += leftLeg
			} else {
				index += rightLeg
			}
			count++
		}
	}
	return result.String()
}