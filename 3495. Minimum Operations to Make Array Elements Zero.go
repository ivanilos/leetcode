func minOperations(queries [][]int) int64 {
    result := int64(0)

    for _, query := range queries {
        // fmt.Println("query = ", solve(query[1]), solve(query[0] - 1))
        result += (solve(query[1]) - solve(query[0] - 1) + 1) / 2 // adding + 1 to round up
    }

    return result
}

func solve(maxVal int) (result int64) {
    nextChunkOps := int64(1)
    for i := int64(4);; i = (i << 2) {
        prevChunkMaxElem := max(0, (i >> 2) - 1)
        curChunkMaxElem := min((i - 1), int64(maxVal))
        total := curChunkMaxElem - prevChunkMaxElem

        // fmt.Println(maxVal, curChunkMaxElem, prevChunkMaxElem, total, nextChunkOps)

        result += total * nextChunkOps
        nextChunkOps++

        if i > int64(maxVal) {
            break
        }
    }

    return result
}
