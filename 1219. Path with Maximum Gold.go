type pair struct {
    x, y int
}

func getMaximumGold(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    paddedGrid := buildPaddedGrid(n, m, grid)

    result := 0
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if paddedGrid[i][j] > 0 {
                visited := make(map[pair]bool)
                
                visited[pair{i, j}] = true
                result = max(result, DFS(i, j, paddedGrid, visited))
                visited[pair{i, j}] = false
            }
        }
    }
    return result
}

func DFS(x, y int, paddedGrid [][]int, visited map[pair]bool) int {
    result := paddedGrid[x][y]
    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx != 0 && dy != 0 {
                continue
            }
            pos := pair{x + dx, y + dy}
            if found, ok := visited[pos]; (!ok || !found) && paddedGrid[x + dx][y + dy] > 0 {
                visited[pair{x + dx, y + dy}] = true
                result = max(result, paddedGrid[x][y] + DFS(x + dx, y + dy, paddedGrid, visited))
                visited[pair{x + dx, y + dy}] = false
            }
        }
    }
    return result
}

func buildPaddedGrid(n, m int, grid [][]int) [][]int {
    paddedGrid := make([][]int, n + 2)

    for i := 0; i < n + 2; i++ {
        paddedGrid[i] = make([]int, m + 2)
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            paddedGrid[i][j] = grid[i - 1][j - 1]
        }
    }
    return paddedGrid
}
