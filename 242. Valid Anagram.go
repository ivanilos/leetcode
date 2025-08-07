func isAnagram(s string, t string) bool {
    freqSum := map[rune]int{}

    for _, c := range s {
        freqSum[c]++
    }

    for _, c := range t {
        freqSum[c]--
    }

    for _, val := range freqSum {
        if val != 0 {
            return false
        }
    }

    return true
}
