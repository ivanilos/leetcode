func appendCharacters(s string, t string) int {
    nextCharInT := 0

    for i := 0; i < len(s); i++ {
        if nextCharInT < len(t) && s[i] == t[nextCharInT] {
            nextCharInT++
        }
    }
    return max(0, len(t) - nextCharInT)
}
