package palindromicsubstrings

func countSubstrings(s string) int {
    res := 0
	
	for i := 0; i < len(s); i++ {
		leftPointer, rightPointer := i, i
		for leftPointer >= 0 && rightPointer < len(s) && s[leftPointer] == s[rightPointer] {
			res++
			leftPointer--
			rightPointer++
		}
	}

	for i := 0; i < len(s) - 1; i++ {
		leftPointer, rightPointer := i, i+1
		for leftPointer >= 0 && rightPointer < len(s) && s[leftPointer] == s[rightPointer] {
			res++
			leftPointer--
			rightPointer++
		}
	}
	return res
}