import "unicode"

func isValid(word string) bool {
    chain := []func(string)bool{hasMinLen, hasAlphanumericOnly, hasVowelAndConsonant}

    for _, f := range chain {
        if !f(word) {
            return false
        }
    }
    return true
}

func hasMinLen(word string) bool {
    return len(word) >= 3
}

func hasAlphanumericOnly(word string) bool {
    for _, r := range word {
        if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func hasVowelAndConsonant(word string) bool {
    vowelsMap := map[rune]bool {
        'a': true,
        'A': true,
        'e': true,
        'E': true,
        'i': true,
        'I': true,
        'o': true,
        'O': true,
        'u': true,
        'U': true,
    }

    vowels := 0
    digits := 0
    for _, r := range word {
        if _, ok := vowelsMap[r]; ok {
            vowels++
        } else if unicode.IsDigit(r) {
            digits++
        }
    }

    return vowels > 0 && len(word) - vowels - digits > 0
}
