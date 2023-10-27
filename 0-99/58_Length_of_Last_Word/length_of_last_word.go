package lengthoflastword

func lengthOfLastWord(s string) int {
	currentLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			currentLength++
			continue
		}

		if currentLength != 0 {
			return currentLength
		}
	}

	return currentLength
}