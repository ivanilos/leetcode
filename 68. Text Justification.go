func fullJustify(words []string, maxWidth int) []string {
    wordsInLine := getWordsInLine(words, maxWidth)

    result := getFormattedWordsInLine(words, wordsInLine, maxWidth)

    return result
}

func getFormattedWordsInLine(words []string, wordsInLine [][]int, maxWidth int) []string {
    result := []string{}

    lines := len(wordsInLine)
    for i := 0; i < lines - 1; i++ {
        result = append(result, getFormattedWordsInFirstLines(wordsInLine[i][0], wordsInLine[i][1], words, maxWidth))
    }

    result = append(result, getFormattedWordsInLastLine(wordsInLine[lines - 1][0], wordsInLine[lines - 1][1], words, maxWidth))

    return result
}

func getFormattedWordsInFirstLines(startIdx, endIdx int, words []string, maxWidth int) string {
    totalTextLen := getTextLen(words, startIdx, endIdx)

    emptySpaces := maxWidth - totalTextLen
    wordGap := emptySpaces / max(1, (endIdx - startIdx))
    extraSpaceGapQt := emptySpaces % max(1, (endIdx - startIdx))

    result := []string{}
    for i := startIdx; i <= endIdx; i++ {
        if i != startIdx {
            curGap := wordGap
            if extraSpaceGapQt > 0 {
                curGap++
                extraSpaceGapQt--
            }
            result = append(result, strings.Repeat(" ", curGap))
        }
        result = append(result, words[i])
    }

    if (startIdx == endIdx) {
        result = append(result, strings.Repeat(" ", emptySpaces))
    }

    return strings.Join(result, "")
}

func getFormattedWordsInLastLine(startIdx, endIdx int, words []string, maxWidth int) string {
    result := []string{}

    textLen := 0
    for i := startIdx; i <= endIdx; i++ {
        if i != startIdx {
            result = append(result, " ")
            textLen++
        }
        result = append(result, words[i])
        textLen += len(words[i])
    }

    result = append(result, strings.Repeat(" ", maxWidth - textLen))

    return strings.Join(result, "")
}

func getTextLen(words []string ,startIdx, endIdx int) int {
    result := 0
    for i := startIdx; i <= endIdx; i++ {
        result += len(words[i])
    }
    return result
}

func getWordsInLine(words []string, maxWidth int) [][]int {
    curLineLen := 0
    curQt := 0
    nextLineFirstIdx := 0

    result := [][]int{}
    for i := 0; i < len(words); i++ {
        if curLineLen + curQt + len(words[i]) <= maxWidth {
            curLineLen += len(words[i])
            curQt++
        } else {
            result = append(result, []int{nextLineFirstIdx, i - 1})
            nextLineFirstIdx = i
            curLineLen = len(words[i])
            curQt = 1
        }
    }
    result = append(result, []int{nextLineFirstIdx, len(words) - 1})

    return result
}
