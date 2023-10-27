package reverseWords

import "strings"

func reverseWords(s string) string {
    var words []string
	lastSymbolIsSpace := true
	end := len(s)
	for i := len(s) - 1; i >=0; i-- {
		if s[i] != ' ' {
			if lastSymbolIsSpace {
				end = i
				lastSymbolIsSpace = false
			}
			continue
		}

		if !lastSymbolIsSpace {
			words = append(words, s[i+1:end+1])
			lastSymbolIsSpace = true
		}
	}

	if !lastSymbolIsSpace {
		words = append(words, s[0:end+1])
	}
	return strings.Join(words," ")
}