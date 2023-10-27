package isomorphicstring

func isIsomorphic(s string, t string) bool {
	hashMapFirst := make(map[byte]byte)
	hashMapSecond := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		valFirst, okFirst := hashMapFirst[s[i]]
		valSecond, okSecond := hashMapSecond[t[i]]

		if !okFirst && !okSecond {
			hashMapFirst[s[i]] = t[i]
			hashMapSecond[t[i]] = s[i]
			continue
		}

		if !okFirst || !okSecond {
			return false
		}

		if valFirst != t[i] || valSecond != s[i] {
			return false
		}
	}
	return true
}