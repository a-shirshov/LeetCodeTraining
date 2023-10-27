package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	masterWord := strs[0]
	prefix := masterWord[0:0]
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < len(masterWord) {
			masterWord = strs[i]
		}
	}

	for i := 1; i <= len(masterWord); i++ {
		prefix = masterWord[0:i]
		for j := 0; j < len(strs); j++ {
			if strs[j][0:i] != prefix {
				return prefix[0:i-1]
			}
		}
	}

	return prefix
}