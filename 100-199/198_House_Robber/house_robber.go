package houserobber

func rob(nums []int) int {
    var res []int
    res = append(res, nums[0])
    for i := 1; i < len(nums); i++ {
        max := res[i-1]
		// for second house: case [2,1]
        if (i - 2) < 0 {
            if nums[i] > max {
                max = nums[i]
            }
            res = append(res, max)
            continue
        }
        possibleTreasure := res[i-2] + nums[i]
        if (possibleTreasure > max) {
            max = possibleTreasure
        }
        res = append(res, max)
    }
    return res[len(res) - 1]
}