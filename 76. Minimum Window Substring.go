func minWindow(s string, t string) string {
    freqT := getFreq(t)
    diffLetters := len(freqT)

    left := 0
    sameFreq := 0
    bestLeft := -1
    bestRight := len(s)

    freqS := map[byte]int{}
    for right := 0; right < len(s); right++ {
        c := s[right]
        freqS[c]++

        if freqS[c] == freqT[c] {
            sameFreq++
        }

        for sameFreq == diffLetters && left <= right {
            if right - left < bestRight - bestLeft {
                bestRight = right
                bestLeft = left
            }

            c := s[left]
            if freqS[c] == freqT[c] {
                sameFreq--
            }
            freqS[c]--
            left++
        }
    }

    if bestLeft == -1 {
        return ""
    }

    return s[bestLeft : bestRight + 1]
}

func getFreq(s string) map[byte]int {
    result := map[byte]int{}

    for _, c := range s {
        result[byte(c)]++
    }

    return result
}
