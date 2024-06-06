func commonChars(words []string) []string {
    freqFirstWord := map[rune]int{}

    for _, c := range(words[0]) {
        freqFirstWord[c]++
    }

    for i := 1; i < len(words); i++ {
        freqCurWord := map[rune]int{}
        for _, c := range(words[i]) {
            freqCurWord[c]++
        }

        for key, _ := range(freqFirstWord) {
            freqFirstWord[key] = min(freqFirstWord[key], freqCurWord[key])
        }
    }

    result := []string{}
    for key, val := range(freqFirstWord) {
        for val > 0 {
            result = append(result, string(key))
            val--
        }
    }
    return result
}
