package romantointeger

func romanToInt(s string) int {
    var hashMap = make(map[byte]int)
	hashMap['I'] = 1
	hashMap['V'] = 5
	hashMap['X'] = 10
	hashMap['L'] = 50
	hashMap['C'] = 100
	hashMap['D'] = 500
	hashMap['M'] = 1000

	result := 0
	for i := 0; i < len(s) - 1; i++ {
		if hashMap[s[i]] >= hashMap[s[i+1]] {
			result += hashMap[s[i]]
			continue
		}
		result -= hashMap[s[i]]
	}
	result += hashMap[s[len(s)-1]]
	return result
}