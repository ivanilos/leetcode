func longestPalindrome(words []string) int {
    matched := 0
    freeStrings := map[string]int{}

    for _, word := range words {
        reversed := string(word[1])
        reversed += string(word[0])
        if val, ok := freeStrings[reversed]; ok && val > 0 {
            matched++
            freeStrings[reversed]--
        } else {
            freeStrings[word]++
        }
    }

    for word, qt := range freeStrings {
        if word[0] == word[1] && qt > 0 {
            return 4 * matched + 2
        }
    }

    return 4 * matched
}
