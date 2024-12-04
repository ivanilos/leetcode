func canMakeSubsequence(str1 string, str2 string) bool {
    nextIdxStr2 := 0

    for i := 0; i < len(str1) && nextIdxStr2 < len(str2); i++ {
        if canMatch(str1[i], str2[nextIdxStr2]) {
            nextIdxStr2++
        }
    }

    return nextIdxStr2 >= len(str2)
}

func canMatch(c1, c2 byte) bool {
    id1 := c1 - 'a'
    id2 := c2 - 'a'

    if id1 == id2 || (id1 + 1) % 26 == id2 {
        return true
    }
    return false
}
