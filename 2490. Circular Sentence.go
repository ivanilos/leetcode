func isCircularSentence(sentence string) bool {
    splitSentence := strings.Split(sentence, " ")

    totalWords := len(splitSentence)
    for i := 0; i < totalWords - 1; i++ {
        prevSz := len(splitSentence[i])
        if splitSentence[i][prevSz - 1] != splitSentence[i + 1][0] {
            return false
        }
    }

    lastWordSz := len(splitSentence[totalWords - 1])
    return splitSentence[0][0] == splitSentence[totalWords - 1][lastWordSz - 1]
}
