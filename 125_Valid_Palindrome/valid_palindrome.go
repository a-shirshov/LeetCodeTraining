package validpalindrome

import (
	"unicode"
)

func isLetterOrDigit(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}

func isPalindrome(s string) bool {
	firstPointer := 0
	secondPointer := len(s) - 1
	runeS := []rune(s)

    for firstPointer < secondPointer {
		left := unicode.ToLower(runeS[firstPointer])
		right := unicode.ToLower(runeS[secondPointer])

		if !isLetterOrDigit(left){
			firstPointer++
			continue
		}

		if !isLetterOrDigit(right) {
			secondPointer--
			continue
		}

		if left != right {
			return false
		}

		firstPointer++
		secondPointer--

	}
	return true
}