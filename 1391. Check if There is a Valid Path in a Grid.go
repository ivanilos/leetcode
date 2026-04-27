var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

const DIRS = 4

var possibleDirs = map[int][]int {
    1: []int{1, 3},
    2: []int{0, 2},
    3: []int{2, 3},
    4: []int{1, 2},
    5: []int{0, 3},
    6: []int{0, 1},
}

func hasValidPath(grid [][]int) bool {
    rows := len(grid)
    cols := len(grid[0])

    visited := make([][]bool, rows)
    for i := 0; i < rows; i++ {
        visited[i] = make([]bool, cols)
    }

    DFS(0, 0, grid, visited)

    return visited[rows - 1][cols - 1]
}

func DFS(x, y int, grid [][]int, visited [][]bool) {
    visited[x][y] = true

    for i := 0; i < DIRS; i++ {
        nx := x + dx[i]
        ny := y + dy[i]

        if isIn(nx, ny, grid) && 
            hasRoad(grid[x][y], i) && 
            hasRoad(grid[nx][ny], (i + 2) % DIRS) &&
            !visited[nx][ny] {
                
            DFS(nx, ny, grid, visited)
        }
    } 
}

func hasRoad(key, dirVal int) bool {
    for _, dir := range possibleDirs[key] {
        if dir == dirVal {
            return true
        }
    }
    return false
}

func isIn(x, y int, grid [][]int) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}
