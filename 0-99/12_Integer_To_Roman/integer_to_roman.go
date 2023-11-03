package integertoroman

import "strings"


var romanLiterals = [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var romanNumbers = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var results strings.Builder
    for i := 0; i < len(romanLiterals); i++ {
		multiplier := num / romanNumbers[i]
		results.WriteString(strings.Repeat(romanLiterals[i], multiplier))
		num -= romanNumbers[i]*multiplier
	}
	return results.String()
}