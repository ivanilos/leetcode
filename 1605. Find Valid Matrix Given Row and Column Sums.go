func restoreMatrix(rowSum []int, colSum []int) [][]int {
    rows := len(rowSum)
    cols := len(colSum)
    result := initMatrix(rows, cols)

    for {
        maxiRow, maxiRowIdx := getMax(rowSum)
        maxiCol, maxiColIdx := getMax(colSum)

        miniVal := min(maxiRow, maxiCol)
        if miniVal == 0 {
            break
        }

        result[maxiRowIdx][maxiColIdx] += miniVal

        rowSum[maxiRowIdx] -= miniVal
        colSum[maxiColIdx] -= miniVal
    }
    return result
}

func initMatrix(rows, cols int) [][]int {
    result := make([][]int, rows)

    for i := 0; i < rows; i++ {
        result[i] = make([]int, cols)
    }
    return result
}

func getMax(v []int) (int, int) {
    result := v[0]
    idx := 0

    for i := 0; i < len(v); i++ {
        if v[i] > result {
            result = v[i]
            idx = i
        }
    }
    return result, idx
}
