var vowels = map[rune]bool{
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
    'A': true,
    'E': true,
    'I': true,
    'O': true,
    'U': true,
}

func reverseVowels(s string) string {
    result := []rune(s)

    left := 0
    right := len(s) - 1
    for left <= right {
        if isVowel(result[left]) && isVowel(result[right]) {
            result[left], result[right] = result[right], result[left]
            left++
            right--
        } else if !isVowel(result[left]) {
            left++
        } else if !isVowel(result[right]) {
            right--
        }
    }

    return string(result)
}

func isVowel(r rune) bool {
    _, ok := vowels[r]
    return ok
}
