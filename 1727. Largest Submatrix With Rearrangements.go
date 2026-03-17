func largestSubmatrix(matrix [][]int) int {
    rows := len(matrix)
    cols := len(matrix[0])

    maxOnesDown := make([][]int, rows)
    for i := 0; i < rows; i++ {
        maxOnesDown[i] = make([]int, cols)
    }

    for j := 0; j < cols; j++ {
        if matrix[rows - 1][j] == 1 {
            maxOnesDown[rows - 1][j] = 1
        }
        
        for i := rows - 2; i >= 0; i-- {
            if matrix[i][j] == 1 {
                maxOnesDown[i][j] = 1 + maxOnesDown[i + 1][j]
            }
        }
    }

    result := 0
    for i := 0; i < rows; i++ {
        result = max(result, solve(i, cols, maxOnesDown))
    }

    return result
}

func solve(curRow int, cols int, maxOnesDown [][]int) int {
    maxRun := make([]int, cols)
    for j := 0; j < cols; j++ {
        maxRun[j] = maxOnesDown[curRow][j]
    }

    slices.Sort(maxRun)
    slices.Reverse(maxRun)

    result := 0
    for i := 0; i < cols; i++ {
        result = max(result, (i + 1) * maxRun[i])
    }

    return result
}
