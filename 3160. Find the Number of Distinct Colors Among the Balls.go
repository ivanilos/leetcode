func queryResults(limit int, queries [][]int) []int {
    colorsFreq := map[int]int{}
    colorsMap := map[int]int{}

    result := []int{}
    curDiff := 0
    for _, query := range queries {
        ball, color := query[0], query[1]

        if _, ok := colorsMap[ball]; ok {
            colorsFreq[colorsMap[ball]]--

            if  colorsFreq[colorsMap[ball]] == 0 {
                curDiff--
            }

            colorsMap[ball] = color
            colorsFreq[color]++

            if colorsFreq[color] == 1 {
                curDiff++
            }
        } else {
            colorsMap[ball] = color
            colorsFreq[color]++

            if colorsFreq[color] == 1 {
                curDiff++
            }
        }

        result = append(result, curDiff)
    }

    return result
}
