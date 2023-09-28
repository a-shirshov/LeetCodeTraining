package dailytemperatures

type tempInfo struct {
	temperature int
	position int
}

func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures))
	var stack []tempInfo
	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 {
			if stack[len(stack) - 1].temperature >= temperatures[i] {
				break
			}

			currentPosition := stack[len(stack) - 1].position
			result[currentPosition] = i - currentPosition
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, tempInfo {
			temperature: temperatures[i], 
			position: i,
		})

	}
	return result
}