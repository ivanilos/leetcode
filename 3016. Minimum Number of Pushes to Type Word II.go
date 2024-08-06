type Pair struct {
    first rune
    second int
}

func minimumPushes(word string) int {
    freq := map[rune]int{}

    for _, c := range(word) {
        freq[c]++
    }

    freqArray := []Pair{}
    for key, val := range(freq) {
        freqArray = append(freqArray, Pair{key, val})
    }

    sort.Slice(freqArray, func(a, b int) bool {
        return freqArray[a].second > freqArray[b].second
    })

    result := 0
    const MAX_ONE = 8
    const MAX_TWO = 16
    const MAX_THREE = 24
    for i := 0; i < len(freqArray); i++ {
        if i < MAX_ONE {
            result += freqArray[i].second
        } else if i < MAX_TWO {
            result += 2 * freqArray[i].second
        } else if i < MAX_THREE {
            result += 3 * freqArray[i].second
        } else {
            result += 4 * freqArray[i].second
        }
    }
    return result
}
