func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])

    result := []int{}
    for col := 0; col < n; col++ {

        diagonalValues := []int{}
        for i, j := 0, col; i < m && j >= 0; i, j = i + 1, j - 1 {
            diagonalValues = append(diagonalValues, mat[i][j])
        }

        if col % 2 == 0 {
            slices.Reverse(diagonalValues)
        }

        result = append(result, diagonalValues...)
    }

    for row := 1; row < m; row++ {
        diagonalValues := []int{}

        for i, j := row, n - 1; i < m && j >= 0; i, j = i + 1, j - 1 {
            diagonalValues = append(diagonalValues, mat[i][j])
        }

        if (row + n) % 2 == 1 {
            slices.Reverse(diagonalValues)
        }

        result = append(result, diagonalValues...)
    }

    return result
}
