func wordSubsets(words1 []string, words2 []string) []string {
    result := []string{}

    words1Freq := getFreq(words1)
    words2Freq := compressFreqs(words2)

    for i := 0; i < len(words1Freq); i++ {
        if isUniversal(words1Freq[i], words2Freq) {
            result = append(result, words1[i])
        }
    }
    return result
}

func isUniversal(freq1 map[int]int, freqUniverse map[int]int) bool {
    return isSubset(freqUniverse, freq1)
}

func isSubset(freq1 map[int]int, freq2 map[int]int) bool {
    for key, val := range freq1 {
        if val > freq2[key] {
            return false
        }
    }
    return true
}

func getFreq(words []string) []map[int]int {
    result := make([]map[int]int, len(words))

    for i := 0; i < len(words); i++ {
        result[i] = getWordFreq(words[i])
    }
    return result
}

func compressFreqs(words []string) map[int]int {
    result := map[int]int{}

    for _, word := range words {
        freq := getWordFreq(word)

        for key, val := range freq {
            result[key] = max(result[key], val)
        }
    }
    return result
}

func getWordFreq(word string) map[int]int {
    result := map[int]int{}

    for _, c := range word {
        idx := int(c - 'a')
        result[idx]++
    }
    return result
}
