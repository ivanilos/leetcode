func numSpecial(mat [][]int) int {
    rowSum := make([]int, len(mat))
    colSum := make([]int, len(mat[0]))

    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            rowSum[i] += mat[i][j]
            colSum[j] += mat[i][j]
        }
    }

    result := 0
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            if mat[i][j] == 1 && rowSum[i] == 1 && colSum[j] == 1 {
                result++
            }
        }
    }

    return result
}
