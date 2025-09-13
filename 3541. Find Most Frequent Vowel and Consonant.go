var vowels = map[rune]bool {
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func maxFreqSum(s string) int {
    freq := map[rune]int{}

    for _, ch := range s {
        freq[ch]++
    }

    maxi := make([]int, 2)
    for key, val := range freq {
        if isVowel(key) {
            maxi[0] = max(maxi[0], val)
        } else {
            maxi[1] = max(maxi[1], val)
        }
    }

    return maxi[0] + maxi[1]
}

func isVowel(ch rune) bool {
    _, ok := vowels[ch]

    return ok
}
