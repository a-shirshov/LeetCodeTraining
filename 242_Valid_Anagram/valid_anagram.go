package isAnagram


func isAnagram(s string, t string) bool {
	if (len(s) != len(t)){
		return false
	}

    var lettersFreq [26]int

	for i:=0; i<len(s); i++{
		lettersFreq[s[i] - 'a']++;
		lettersFreq[t[i] - 'a']--;
	}
	
	for _, letterFreq := range lettersFreq {
		if(letterFreq != 0) {
			return false
		}
	}

	return true
}