package text_just

import (
	"bytes"
	"io"
	"strings"
)

func printWordsLast(sliceOfWords []string, wordsLenCount int, maxWidth int, w io.Writer) {
	count := 0
	w.Write([]byte(sliceOfWords[0]))
	count += len(sliceOfWords[0])
	for i := 1; i < len(sliceOfWords); i++ {
		w.Write([]byte(" "))
		w.Write([]byte(sliceOfWords[i]))
		count += len(sliceOfWords[i]) + 1
	}

	for i := count; i < maxWidth; i++ {
		w.Write([]byte(" "))
	}
}

func printWords(sliceOfWords []string, wordsLenCount int, maxWidth int, w io.Writer) {
	if len(sliceOfWords) == 1 {
		w.Write([]byte(sliceOfWords[0]))
		for i := len(sliceOfWords[0]); i < maxWidth; i++ {
			w.Write([]byte(" "))
		}
		w.Write([]byte("\n"))
		return
	}
	amountOfWordNeededSpaces := len(sliceOfWords) - 1
	spaceForSpaces := maxWidth - wordsLenCount
	amountOfShouldBeSpaces := spaceForSpaces / amountOfWordNeededSpaces
	amountOfLeftExSpaces := spaceForSpaces % amountOfWordNeededSpaces
	for i := 0; i < len(sliceOfWords); i++ {
		if i == 0 {
			w.Write([]byte(sliceOfWords[i]))
			continue
		}

		if i == len(sliceOfWords) - 1 {
			for j := 0; j < amountOfShouldBeSpaces; j++ {
				w.Write([]byte(" "))
			}
		} else {
			for j := 0; j < amountOfShouldBeSpaces; j++ {
				w.Write([]byte(" "))
			}
			if amountOfLeftExSpaces != 0 {
				w.Write([]byte(" "))
				amountOfLeftExSpaces--
			}
		}
		w.Write([]byte(sliceOfWords[i]))
	}
	w.Write([]byte("\n"))
}

func fullJustify(words []string, maxWidth int) []string {
    var sliceOfWords []string
	buf := &bytes.Buffer{}
	wordsLenCount := 0
	wordsLenCountWithSpaces := 0
	for i := 0; i < len(words); i++ {
		if len(sliceOfWords) == 0 {
			sliceOfWords = append(sliceOfWords, words[i])
			wordsLenCount += len(words[i])
			wordsLenCountWithSpaces = wordsLenCount + len(sliceOfWords) - 1
			continue
		}

		if wordsLenCountWithSpaces + len(words[i]) + 1 > maxWidth {
			printWords(sliceOfWords, wordsLenCount, maxWidth, buf)
			sliceOfWords = []string{}
			wordsLenCount = 0
			wordsLenCountWithSpaces = 0
			i--
		} else {
			sliceOfWords = append(sliceOfWords, words[i])
			wordsLenCount += len(words[i])
			wordsLenCountWithSpaces = wordsLenCount + len(sliceOfWords) - 1
		}
	}

	if len(sliceOfWords) != 0 {
		printWordsLast(sliceOfWords, wordsLenCount, maxWidth, buf)
	}
	return strings.Split(buf.String(), "\n")
}