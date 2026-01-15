func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
    maxContinuous := min(getMaxContinuous(vBars), getMaxContinuous(hBars))

    return (1 + maxContinuous) * (1 + maxContinuous)
}

func getMaxContinuous(bars []int) int {
    slices.Sort(bars)

    last := bars[0] - 2
    cur := 0
    result := 0
    for i := 0; i < len(bars); i++ {
        if bars[i] == last + 1 {
            cur++
        } else {
            cur = 1
        }
        result = max(result, cur)
        last = bars[i]
    }

    return result
}
