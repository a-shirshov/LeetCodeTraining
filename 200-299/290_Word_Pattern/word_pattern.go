package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	var sb strings.Builder
	patternLength := len(pattern)
    hashMap := make(map[byte]string)
	hashMapWordUsed := make(map[string]struct{})
	var i, j int
	for i, j = 0, 0; i < len(s) && j < patternLength; i++ {
		if s[i] != ' ' {
			sb.WriteByte(s[i])
			continue
		}

		word, ok := hashMap[pattern[j]]
		if !ok {
			_, isUsed := hashMapWordUsed[sb.String()]
			if isUsed {
				return false
			}
			hashMap[pattern[j]] = sb.String()
			hashMapWordUsed[sb.String()] = struct{}{}
		} else if word != sb.String() {
			return false
		}
		j++
		sb.Reset()
	}

	if i == len(s) && j == patternLength - 1 {
		word, ok := hashMap[pattern[j]]
		if !ok {
			_, isUsed := hashMapWordUsed[sb.String()]
			if isUsed {
				return false
			}
			hashMap[pattern[j]] = sb.String()
			hashMapWordUsed[sb.String()] = struct{}{}
		} else if word != sb.String() {
			return false
		}
		return true
	}

	return false
}