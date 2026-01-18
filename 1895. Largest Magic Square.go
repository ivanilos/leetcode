func largestMagicSquare(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    for side := min(rows, cols); side >= 2; side-- {
        for i := 0; i <= rows - side; i++ {
            for j := 0; j <= cols - side; j++ {
                if isMagical(i, j, side, grid) {
                    return side
                }
            }
        }
    }

    return 1
}

func isMagical(topX, topY, side int, grid [][]int) bool {
    rowSum := make([]int, side)
    colSum := make([]int, side)
    diag := make([]int, 2)

    for i := topX; i < topX + side; i++ {
        for j := topY; j < topY + side; j++ {
            rowSum[i - topX] += grid[i][j]
            colSum[j - topY] += grid[i][j]

            if (i - topX) == (j - topY) {
                diag[0] += grid[i][j]
            }
            if (i - topX) + (j - topY) == side - 1 {
                diag[1] += grid[i][j]
            }
        }
    }

    if diag[0] != diag[1] {
        return false
    }
    for i := 1; i < side; i++ {
        if rowSum[i] != rowSum[i - 1] {
            return false
        }
        if colSum[i] != colSum[i - 1] {
            return false
        }
    }

    return rowSum[0] == colSum[0] && colSum[0] == diag[0]
}
