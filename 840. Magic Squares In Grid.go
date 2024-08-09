func numMagicSquaresInside(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    result := 0
    for leftX := 0; leftX + 2 < rows; leftX++ {
        for topY := 0; topY + 2 < cols; topY++ {
            if (isMagicSquare(grid, leftX, topY)) {
                result++
            }
        }
    }
    return result
}

func isMagicSquare(grid [][]int, leftX, topY int) bool {
    sums := map[int]int{}
    nums := map[int]int{}

    rowSum := []int{0, 0, 0}
    colSum := []int{0, 0, 0}
    diagSum := []int{0, 0}

    for row := leftX; row < leftX + 3; row++ {
        for col := topY; col < topY + 3; col++ {
            if grid[row][col] <= 0 || grid[row][col] > 9 {
                return false
            } 
            nums[grid[row][col]]++

            if (nums[grid[row][col]] > 1) {
                return false
            }

            rowSum[row - leftX] += grid[row][col]
            colSum[col - topY] += grid[row][col]
            if (row - leftX == col - topY) {
                diagSum[0] += grid[row][col]
            }
            if (row - leftX + col - topY == 2) {
                diagSum[1] += grid[row][col]
            }
        }
    }

    for i := 0; i < 3; i++ {
        sums[rowSum[i]]++
        sums[colSum[i]]++
        if i < 2 {
            sums[diagSum[i]]++
        }
    }

    return len(sums) == 1
}
