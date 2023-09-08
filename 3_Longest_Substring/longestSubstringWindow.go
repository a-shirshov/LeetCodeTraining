package longestSubstring

import (

)

func max(first, second int) int {
	if (first > second) {
		return first
	}
	return second
}

func LengthOfLongestSubstringWindow(s string) int {
	length := len(s)
	start, result := 0, 0
	charIndexMap := make(map[byte]int)
	for end := 0; end != length; end++ {
		index, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			result = max(result, end - start)
			for i := start; i < index; i ++ {
				delete(charIndexMap,s[i])
			}
			start = index + 1
		}
		charIndexMap[s[end]] = end
	}
	result = max(result, length - start)
	return result;
}