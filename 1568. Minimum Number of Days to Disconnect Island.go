func minDays(grid [][]int) int {
    components := countComponents(grid, -1, -1)

    fmt.Println(components)
    if components > 1 || components == 0 {
        return 0
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if (grid[i][j] == 1) {
                components = countComponents(grid, i, j)
                if components > 1 || components == 0 {
                    return 1
                }
            }
        }
    }
    return 2
}

func countComponents(grid [][]int, x, y int) int {
    used := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        used[i] = make([]int, len(grid[i]))
    }

    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 && used[i][j] == 0 && !(x == i && y == j) {
                result++
                DFS(grid, used, i, j, x, y)
            }
        }
    }
    return result
}

func DFS(grid [][]int, used [][]int, sx, sy, notX, notY int) {
    dx := []int{1, 0, -1, 0}
    dy := []int{0, 1, 0, -1}

    used[sx][sy] = 1

    for i := 0; i < 4; i++ {
        nx := sx + dx[i]
        ny := sy + dy[i]

        if isIn(grid, nx, ny) && grid[nx][ny] == 1 && used[nx][ny] == 0 && !(nx == notX && ny == notY) {
            DFS(grid, used, nx, ny, notX, notY)
        }
    }
}

func isIn(grid [][]int, nx, ny int) bool {
    return 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0])
}
