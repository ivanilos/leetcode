func appealSum(s string) int64 {
    pos := map[rune][]int{}

    result := int64(0)
    for i, r := range s {
        pos[r] = append(pos[r], i)

        result += getScore(pos)
    }

    return result
}

func getScore(pos map[rune][]int) int64 {
    vals := []int64{}

    for _, valList := range pos {
        val := valList[len(valList) - 1]
        vals = append(vals, int64(val))
    }

    slices.Sort(vals)

    result := int64(0)
    curAdd := int64(1)
    lastIdx := vals[len(vals) - 1]
    for i := len(vals) - 2; i >= 0; i-- {
        result += (lastIdx - vals[i]) * curAdd
        lastIdx = vals[i]
        curAdd++
    }

    result += (lastIdx + 1) * curAdd

    return result
}
