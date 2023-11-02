package isSubsequence

func isSubsequence(s string, t string) bool {
    index := 0
	for i := 0; i < len(t) && index < len(s); i++ {
		if t[i] == s[index] {
			index++
		}
	}

	return index == len(s)
}