import "unicode"

func isPalindrome(s string) bool {
    runes := []rune(s)

    left := 0
    right := len(runes) - 1

    for left <= right {
        for left < len(runes) && !isAlphanumeric(runes[left]) {
            left++
        }
        for right >= 0 && !isAlphanumeric(runes[right]) {
            right--
        }

        if left < len(runes) && right >= 0 && !palindromeEqual(runes[left], runes[right]) {
            return false
        }
        left++
        right--
    }

    return true
}

func isAlphanumeric(r rune) bool {
    return unicode.IsDigit(r) || unicode.IsLetter(r)
}

func palindromeEqual(a, b rune) bool {
    if unicode.IsDigit(a) && unicode.IsDigit(b) {
        return a == b
    }
    if unicode.IsLetter(a) && unicode.IsLetter(b) {
        return unicode.ToLower(a) == unicode.ToLower(b)
    }
    return false
}
