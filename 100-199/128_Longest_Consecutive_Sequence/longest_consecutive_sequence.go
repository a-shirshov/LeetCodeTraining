package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]bool)
    for _,value := range nums {
		hashMap[value] = true
	}

	max := 0
	for _,value := range nums {
		if hashMap[value-1] {
			continue
		}

		length := 1
		for hashMap[value+length]{
			length++
		}
		if (length > max){
			max = length
		}
	}
	return max
}