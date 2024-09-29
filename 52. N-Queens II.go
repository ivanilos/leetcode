var result int = 0

func totalNQueens(n int) int {
    usedCol := make([]bool, n)
    usedDiag := map[int]bool{}
    usedSecDiag := map[int]bool{}
    result = 0

    solve(0, n, usedCol, usedDiag, usedSecDiag)

    return result
}

func solve(row, n int, usedCol []bool, usedDiag map[int]bool, usedSecDiag map[int]bool) {
    if row >= n {
        result++
        return
    }

    for i := 0; i < n; i++ {
        if isValidPlacement(row, i, usedCol, usedDiag, usedSecDiag) {
            usedCol[i] = true
            usedDiag[row - i] = true
            usedSecDiag[i + row] = true

            solve(row + 1, n, usedCol, usedDiag, usedSecDiag)

            usedCol[i] = false
            usedDiag[row - i] = false
            usedSecDiag[i + row] = false
        }
    }
}

func isValidPlacement(row, col int, usedCol []bool, usedDiag map[int]bool, usedSecDiag map[int]bool) bool {
    if usedCol[col] {
        return false
    }
    if usedDiag[row - col] {
        return false
    }
    if usedSecDiag[row + col] {
        return false
    }
    return true
}
