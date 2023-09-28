package longestSubstring

import (

)

func LengthOfLongestSubstring(s string) int {
	max_length := 0
	length := 0
	hashTable := make(map[byte]int)

	if len(s) == 1 {
		return 1
	}

	for i:= 0; i < len(s) - 1; i++ {
		for k := range hashTable {
			delete(hashTable, k)
		}
		length = 0
		hashTable[s[i]] = i;
		length++
		for j := i + 1; j < len(s); j++ {
			_, ok := hashTable[s[j]]
			if ok {
				if max_length < length {
					max_length = length
				}
				length = 0;
				for k := range hashTable {
					delete(hashTable, k)
				}
				break
			}
			hashTable[s[j]] = j
			length++
		}

		if max_length < length {
			max_length = length
		}
	}

	return max_length
}