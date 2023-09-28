package groupanagrams

import (
	"sort"
	"strings"
)

func isAnagram(str1 string, str2 string) bool {
	if (len(str1) != len(str2)){
		return false
	}

	var letterFreq [26]int
	for i:=0; i<len(str1); i++ {
		letterFreq[str1[i]-'a']++
		letterFreq[str2[i]-'a']--
	}

	for _, freq := range letterFreq {
		if freq != 0 {
			return false;
		}
	}
	
	return true;
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	
	for i:=0; i<len(strs); i++{
		sortedStr := sortString(strs[i])
		if _, ok := hashMap[sortedStr]; !ok {
			hashMap[sortedStr] = []string{}
		}
		hashMap[sortedStr] = append(hashMap[sortedStr], strs[i])
	}

	result := make([][]string, len(hashMap))
	i := 0
	for _, value := range hashMap{
		result[i] = value
		i++
	}
	return result
}

func sortString(str string) string{
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr,"")
}