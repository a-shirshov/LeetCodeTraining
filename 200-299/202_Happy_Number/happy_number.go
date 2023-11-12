package happynumber

func isHappy(n int) bool {
	hashMap:= make(map[int]struct{})
    getNewNumber := func(n int) int {
		sum := 0
		for n != 0{
			digit := n % 10
			sum += digit * digit
			n = n / 10
		}
		return sum
	}

	for n != 1 {
		_, was := hashMap[n]
		if was {
			return false
		}
		hashMap[n] = struct{}{}
		n = getNewNumber(n)
	}

	return true
}