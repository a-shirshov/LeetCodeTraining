package maximumaveragesubarrayi

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
   var windowSum int
   start := 0
   max := math.Inf(-1)

   for end := 0; end < len(nums); end++ {
        windowSum += nums[end]

        if (end - start + 1) == k{
            max = math.Max(max, float64(windowSum))
            windowSum -= nums[start]
            start++
        }
    }
    return max/float64(k)
}