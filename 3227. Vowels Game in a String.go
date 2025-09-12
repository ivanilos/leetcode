var vowels = map[rune]bool{
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func doesAliceWin(s string) bool {
    if hasVowel(s) {
        return true
    } else {
        return false
    }
}

func hasVowel(s string) bool {
    for _, ch := range s {
        if _, ok := vowels[ch]; ok {
            return true
        }
    }
    return false
}
