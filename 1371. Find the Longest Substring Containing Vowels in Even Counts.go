func findTheLongestSubstring(s string) int {
    lastPos := map[int]int{}
    id := map[byte]int{'a' : 0, 'e' : 1, 'i': 2, 'o': 3, 'u': 4}

    lastPos[0] = -1
    curMask := 0
    result := 0
    for i := 0; i < len(s); i++ {
        if isVowel(s[i], id) {
            bit := id[s[i]]
            curMask ^= (1 << bit)
        }
        if _, ok := lastPos[curMask]; ok {
            result = max(result, i - lastPos[curMask])
        } else {
            lastPos[curMask] = i
        }
    }
    return result
}

func isVowel(c byte, vowels map[byte]int) bool {
    _, ok := vowels[c]
    return ok
}
