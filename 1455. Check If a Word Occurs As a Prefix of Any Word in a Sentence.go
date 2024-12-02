func isPrefixOfWord(sentence string, searchWord string) int {
    splittedSentence := strings.Split(sentence, " ")

    for idx, word := range splittedSentence {
        if len(word) >= len(searchWord) && searchWord == word[: len(searchWord)] {
            return idx + 1
        }
    }
    return -1
}
