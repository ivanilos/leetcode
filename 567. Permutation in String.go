func checkInclusion(s1 string, s2 string) bool {
    freqS1 := getMapFreq(s1)
    diffLetters := len(freqS1)

    curFreqS2 := map[byte]int{}
    curFreqDone := 0
    l := 0
    for r := 0; r < len(s2); r++ {
        nextChar := s2[r]
        curFreqS2[nextChar]++

        if _, ok := freqS1[nextChar]; ok && curFreqS2[nextChar] == freqS1[nextChar] {
            curFreqDone++
        }
        if r - l + 1 > len(s1) {
            leftmostLetter := s2[l]
            if _, ok := freqS1[leftmostLetter]; ok && curFreqS2[leftmostLetter] == freqS1[leftmostLetter] {
                curFreqDone--
            }
            curFreqS2[leftmostLetter]--
            l++
        }
        if curFreqDone == diffLetters && r - l + 1 == len(s1) {
            return true
        }
    }
    return false
}

func getMapFreq(s string) map[byte]int {
    result := map[byte]int{}

    for _, c := range s {
        result[byte(c)]++
    }
    return result
}
