func maxMatrixSum(matrix [][]int) int64 {
    result := int64(0)

    minAbsolute := abs(matrix[0][0])
    negativeQt := 0
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            result += abs(matrix[i][j])
            minAbsolute = min(minAbsolute, abs(matrix[i][j]))

            if matrix[i][j] < 0 {
                negativeQt++
            }

        }
    }

    if negativeQt % 2 == 1 {
        return result - 2 * minAbsolute
    } else {
        return result
    }
}

func abs(x int) int64 {
    if x < 0 {
        return int64(-x)
    }
    return int64(x)
}
