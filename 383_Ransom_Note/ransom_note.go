package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
    hashMap := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		hashMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		v := hashMap[ransomNote[i]]
		if v == 0 {
			return false
		}
		hashMap[ransomNote[i]]--
	}
	return true
}