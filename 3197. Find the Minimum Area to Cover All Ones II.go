const INF = int(1e9)

func minimumSum(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    result := INF

    // horizontal split
    for i := 0; i < rows; i++ {
        result = min(result, 
                getMinimumSumWithRectangles(grid, 0, i, 0, cols - 1, 1) + getMinimumSumWithRectangles(grid, i + 1, rows - 1, 0, cols - 1, 2))

        result = min(result, 
                getMinimumSumWithRectangles(grid, 0, i, 0, cols - 1, 2) + getMinimumSumWithRectangles(grid, i + 1, rows - 1, 0, cols - 1, 1))
    }

    // vertical split
    for j := 0; j < cols; j++ {
        result = min(result, 
                getMinimumSumWithRectangles(grid, 0, rows - 1, 0, j, 1) + getMinimumSumWithRectangles(grid, 0, rows - 1, j + 1, cols - 1, 2))

        result = min(result, 
                getMinimumSumWithRectangles(grid, 0, rows - 1, 0, j, 2) + getMinimumSumWithRectangles(grid, 0, rows - 1, j + 1, cols - 1, 1))
    }

    return result
}

func getMinimumSumWithRectangles(grid [][]int,
        xmin, xmax, ymin, ymax,  maxQuantity int) int {

    xminWithOne := INF
    xmaxWithOne := 0
    yminWithOne := INF
    ymaxWithOne := 0
    if maxQuantity == 1 {
        for i := xmin; i <= xmax; i++ {
            for j := ymin; j <= ymax; j++ {
                if grid[i][j] == 1 {
                    xminWithOne = min(xminWithOne, i)
                    xmaxWithOne = max(xmaxWithOne, i)
                    yminWithOne = min(yminWithOne, j)
                    ymaxWithOne = max(ymaxWithOne, j)
                }
            }
        }

        if xmaxWithOne < xminWithOne || ymaxWithOne < yminWithOne {
            return 0
        }

        return (xmaxWithOne - xminWithOne + 1) * (ymaxWithOne - yminWithOne + 1)

    } else {
        result := INF
        // horizontal split
        for i := xmin; i <= xmax; i++ {
            result = min(result, 
                    getMinimumSumWithRectangles(grid, xmin, i, ymin, ymax, 1) + getMinimumSumWithRectangles(grid, i + 1, xmax, ymin, ymax, 1))
        }

        // vertical split
        for j := ymin; j <= ymax; j++ {
            result = min(result, 
                    getMinimumSumWithRectangles(grid, xmin, xmax, ymin, j, 1) + getMinimumSumWithRectangles(grid, xmin, xmax, j + 1, ymax, 1))
        }

        return result
    }
}
