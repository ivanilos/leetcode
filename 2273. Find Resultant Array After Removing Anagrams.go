func removeAnagrams(words []string) []string {
    wordToSorted := map[string]string{}

    for _, word := range words {
        sorted := []byte(word)
        slices.Sort(sorted)
        wordToSorted[word] = string(sorted)
    }

    st := []string{}
    for _, word := range words {
        if len(st) > 0 {
            topSorted := wordToSorted[st[len(st) - 1]]
            if topSorted != wordToSorted[word] {
                st = append(st, word)
            }
        } else {
            st = append(st, word)
        }
    }

    return st
}
