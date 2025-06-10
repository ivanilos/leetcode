func maxDifference(s string) int {
    freq := map[rune]int{}

    for _, c := range s {
        freq[c]++
    }

    const INF = int(1e9)

    oddFreq := 0
    evenFreq := INF

    for _, val := range freq {
        if val % 2 == 0 {
            evenFreq = min(evenFreq, val)
        } else {
            oddFreq = max(oddFreq, val)
        }
    }

    return oddFreq - evenFreq
}
