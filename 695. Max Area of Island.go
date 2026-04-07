var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func maxAreaOfIsland(grid [][]int) int {
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                result = max(result, DFS(i, j, grid))
            }
        }
    }

    return result
}

func DFS(x, y int, grid [][]int) int {
    result := 1

    grid[x][y] = 0

    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]

        if isIn(nx, ny, grid) && grid[nx][ny] == 1 {
            result += DFS(nx, ny, grid)
        }
    }

    return result
}

func isIn(x, y int, grid [][]int) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}
