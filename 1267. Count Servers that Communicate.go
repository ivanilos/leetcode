var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func countServers(grid [][]int) int {
    result := 0

    rows := make([]int, len(grid))
    cols := make([]int, len(grid[0]))
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            rows[i] += grid[i][j]
            cols[j] += grid[i][j]
        }
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
                result++
            }
        }
    }
    return result
}
