const MOD = 12345

func constructProductMatrix(grid [][]int) [][]int {
    rows := len(grid)
    cols := len(grid[0])

    rowProdPrefix, rowProdSuffix := getRowsProdPrefixAndSuffix(rows, cols, grid)
    allRowsProdPrefix, allRowsProdSuffix := getAllRowsProdPrefixAndSuffix(rows, cols, grid, rowProdPrefix, rowProdSuffix)

    result := make([][]int, rows)
    for i := 0; i < rows; i++ {
        result[i] = make([]int, cols)

        for j := 0; j < cols; j++ {
            result[i][j] = getResult(i, j, rows, cols, allRowsProdPrefix, allRowsProdSuffix, rowProdPrefix, rowProdSuffix)
        }
    }

    return result
}

func getRowsProdPrefixAndSuffix(rows, cols int, grid [][]int) ([][]int, [][]int) {
    rowProdPrefix := make([][]int, rows)
    rowProdSuffix := make([][]int, rows)
    for i := 0; i < rows; i++ {
        rowProdPrefix[i] = make([]int, cols)
        rowProdSuffix[i] = make([]int, cols)
    }

    for i := 0; i < rows; i++ {
        rowProdPrefix[i][0] = grid[i][0] % MOD
        for j := 1; j < cols; j++ {
            rowProdPrefix[i][j] = ((rowProdPrefix[i][j - 1] % MOD) * grid[i][j]) % MOD
        }

        rowProdSuffix[i][cols - 1] = grid[i][cols - 1] % MOD
        for j := cols - 2; j >= 0; j-- {
            rowProdSuffix[i][j] = ((rowProdSuffix[i][j + 1] % MOD) * grid[i][j]) % MOD
        }
    }

    return rowProdPrefix, rowProdSuffix
}

func getAllRowsProdPrefixAndSuffix(rows, cols int, grid [][]int, rowProdPrefix, rowProdSuffix [][]int) ([]int, []int) {
    allRowsProdPrefix := make([]int, rows)
    allRowsProdSuffix := make([]int, rows)

    allRowsProdPrefix[0] = rowProdPrefix[0][cols - 1]
    for i := 1; i < rows; i++ {
        allRowsProdPrefix[i] = (allRowsProdPrefix[i - 1] * rowProdPrefix[i][cols - 1]) % MOD
    }

    allRowsProdSuffix[rows - 1] = rowProdSuffix[rows - 1][0]
    for i := rows - 2; i >= 0; i-- {
        allRowsProdSuffix[i] = (allRowsProdSuffix[i + 1] * rowProdSuffix[i][0]) % MOD
    }

    return allRowsProdPrefix, allRowsProdSuffix
}

func getResult(i, j, rows, cols int, allRowsProdPrefix, allRowsProdSuffix []int, rowProdPrefix, rowProdSuffix [][]int) int {
    result := 1

    if i - 1 >= 0 {
        result = (result * allRowsProdPrefix[i - 1]) % MOD
    }

    if i + 1 < rows {
        result = (result * allRowsProdSuffix[i + 1]) % MOD
    }

    if j - 1 >= 0 {
        result = (result * rowProdPrefix[i][j - 1]) % MOD
    }

    if j + 1 < cols {
        result = (result * rowProdSuffix[i][j + 1]) % MOD
    }

    return result
}
