func canBeTypedWords(text string, brokenLetters string) int {
    brokenLettersMap := map[rune]bool{}
    for _, ch := range brokenLetters {
        brokenLettersMap[ch] = true
    }

    words := strings.Split(text, " ")

    result := len(words)
    for _, word := range words {
        for _, ch := range word {
            if _, ok := brokenLettersMap[ch]; ok {
                result--
                break
            }
        }
    }

    return result
}
