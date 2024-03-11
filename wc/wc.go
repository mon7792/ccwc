package wc

// CharacterCount counts the number of characters in the input
func CharacterCount(inCh <-chan []byte) int {
	var charCount = 0
	for inp := range inCh {
		charCount += charCountSlice(inp)
	}
	return charCount
}

func charCountSlice(inp []byte) int {
	size := len(inp)
	charCount := size
	if inp[size-1] == 0 {
		for i := size - 1; i >= 0; i-- {
			if inp[i] == 0 {
				charCount = charCount - 1
			}
		}
	}
	return charCount
}

// LineCount counts the number of lines in the input
func LineCount(inCh <-chan []byte) int {
	var lineCount = 0
	for inp := range inCh {
		lineCount += lineCountSlice(inp)
	}
	return lineCount
}

func lineCountSlice(inp []byte) int {
	size := len(inp)
	lineCount := 0
	for i := 0; i < size; i++ {
		if inp[i] == 10 {
			lineCount++
		}
		if inp[i] == 0 {
			break
		}
	}

	return lineCount
}

// WordCount counts the number of words in the input
func WordCount(inCh <-chan []byte) int {
	var wordCount = 0
	var partialW = []byte{}
	for inp := range inCh {
		wdcnt := 0
		wdcnt, partialW = wordCountSlice(inp, partialW)
		wordCount += wdcnt
	}
	return wordCount - 1
}

func wordCountSlice(inp []byte, partialW []byte) (int, []byte) {
	size := len(inp)
	wordCount := 0
	var wrdDetected bool
	for i := 0; i < size; i++ {
		if inp[i] == 32 || inp[i] == 10 || inp[i] == 9 || inp[i] == 13 || inp[i] == 12 {
			wrdDetected = false
			partialW = []byte{}
		} else {
			if !wrdDetected && len(partialW) == 0 {
				wrdDetected = true
				wordCount++
			}
			partialW = append(partialW, inp[i])
		}
	}
	return wordCount, partialW
}
