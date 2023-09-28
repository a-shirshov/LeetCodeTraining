package fruitintobasket


func totalFruit(fruits []int) int {
    treeTypesMap := make(map[int]bool)
	var max int
	var start int
	for end, treeType := range fruits {
		if len(treeTypesMap) < 2 || treeTypesMap[treeType] {
			treeTypesMap[treeType] = true
			if (end - start + 1) > max {
				max = end - start + 1
			}
			continue
		}

		treeTypesMap = make(map[int]bool)
		treeTypesMap[fruits[end-1]] = true
		treeTypesMap[treeType] = true
		start = end - 1

		for fruits[start] == fruits[start-1]{
			start --
		}
	}
	return max
}