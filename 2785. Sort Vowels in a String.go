var vowelsMap = map[rune]bool{
    'a' : true,
    'e' : true,
    'i' : true,
    'o' : true,
    'u' : true,
    'A' : true,
    'E' : true,
    'I' : true,
    'O' : true,
    'U' : true,
}

func sortVowels(s string) string {
    result := []rune(s)

    vowels := []rune{}
    for _, ch := range s {
        if isVowel(ch) {
            vowels = append(vowels, ch)
        }
    }

    slices.SortFunc(vowels, func(a, b rune) int {
        return int(a) - int(b)
    })

    nextVowel := 0
    for i, ch := range result {
        if isVowel(ch) {
            result[i] = vowels[nextVowel]
            nextVowel++
        }
    }

    return string(result)
}

func isVowel(ch rune) bool {
    _, ok := vowelsMap[ch]

    return ok
}
