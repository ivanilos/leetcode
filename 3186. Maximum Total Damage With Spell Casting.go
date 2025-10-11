func maximumTotalDamage(power []int) int64 {
    slices.Sort(power)

    freqArr := getFreqArr(power)
    
    dp := make([]int64, len(freqArr))

    result := solve(0, freqArr, dp)
    return result

}

func getFreqArr(power []int) [][]int {
    freqArr := [][]int{}

    cur := -1
    qt := 0
    for _, p := range power {
        if p == cur {
            qt++
        } else {
            if cur != -1 {
                freqArr = append(freqArr, []int{cur, qt})
            }
            cur = p
            qt = 1
        }
    }

    freqArr = append(freqArr, []int{cur, qt})

    return freqArr
}

func solve(pos int, freqArr [][]int, dp []int64) int64 {
    if pos >= len(freqArr) {
        return 0
    }

    if dp[pos] != 0 {
        return dp[pos]
    }

    // not take spell
    result := solve(pos + 1, freqArr, dp)

    // take spell
    nextPos := pos + 1
    for nextPos < len(freqArr) && freqArr[nextPos][0] <= freqArr[pos][0] + 2 {
        nextPos++
    }
    score := int64(freqArr[pos][0]) * int64(freqArr[pos][1])
    result = max(result, score + solve(nextPos, freqArr, dp))

    dp[pos] = result
    return result
}
