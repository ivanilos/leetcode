const WHITE = 0
const GRAY = 1
const BLACK = 2

var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func containsCycle(grid [][]byte) bool {
    color := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        color[i] = make([]int, len(grid[i]))
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if color[i][j] == WHITE {
                hasCycle := DFS(i, j, -1, -1, grid[i][j], grid, color)
                if hasCycle {
                    return true
                }
            }
        }
    }

    return false
}

func DFS(x, y, parX, parY int, val byte, grid [][]byte, color [][]int) bool {
    color[x][y] = GRAY

    result := false
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]

        if isIn(nx, ny, grid) && !(nx == parX && ny == parY) {
            if grid[nx][ny] == val && color[nx][ny] == GRAY {
                return true
            } else if grid[nx][ny] == val && color[nx][ny] == WHITE {
                result = result || DFS(nx, ny, x, y, val, grid, color)
            }
        }
    }

    color[x][y] = BLACK

    return result
}

func isIn(x, y int, grid [][]byte) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}
