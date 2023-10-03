package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	resultIndex := -1
	for i := 0; i <= len(haystack) - len(needle); i++ {
		if haystack[i] == needle[0] {
			resultIndex = i;
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					resultIndex = -1
					break
				}
			}
			
			if resultIndex != -1 {
				return resultIndex
			}
		}
	}
	return resultIndex
}