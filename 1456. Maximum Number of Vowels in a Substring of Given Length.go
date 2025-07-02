var vowels = map[byte]bool {
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func maxVowels(s string, k int) int {
    left := 0
    right := 0

    curVowels := 0
    for ; right < k; right++ {
        if isVowel(s[right]) {
            curVowels++
        }
    }

    result := curVowels
    for ; right < len(s); left, right = left + 1, right + 1 {
        if isVowel(s[right]) {
            curVowels++
        }
        if isVowel(s[left]) {
            curVowels--
        }

        result = max(result, curVowels)
    }

    return result
}

func isVowel(r byte) bool {
    _, ok := vowels[r]
    return ok
}
