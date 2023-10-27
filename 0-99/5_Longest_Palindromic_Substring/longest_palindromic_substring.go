package longestpalindromicsubstring

func expandFromCenter(s string, left int, right int) int {
	sLength := len(s)
	for(left >=0 && right < sLength && s[left] == s[right]) {
		left--
		right++
	}
	return right - left - 1
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func longestPalindrome(s string) string {
	start, end := 0, 0
    for i := 0; i < len(s); i++ {
		len1 := expandFromCenter(s, i, i)
		len2 := expandFromCenter(s, i, i + 1)
		len := max(len1,len2)
		if (len > end - start) {
			start = i - (len - 1) / 2
			end = i + len / 2
		}
	}
	return s[start:end + 1]
}