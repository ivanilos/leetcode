func isSubsequence(s string, t string) bool {
    nextSIdx := 0

    for i := 0; i < len(t) && nextSIdx < len(s); i++ {
        if t[i] == s[nextSIdx] {
            nextSIdx++
        }
    }
    return nextSIdx >= len(s)
}
