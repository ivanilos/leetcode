var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

var INF = int(1e9)

func nearestExit(maze [][]byte, entrance []int) int {
    return BFS(maze, entrance)
}

func BFS(maze [][]byte, entrance []int) int {
    n := len(maze)
    m := len(maze[0])

    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, m)
        for j := 0; j < m; j++ {
            dist[i][j] = INF
        }
    }

    dist[entrance[0]][entrance[1]] = 0
    start := []int{entrance[0], entrance[1]}

    q := [][]int{start}
    for len(q) > 0 {
        curx := q[0][0]
        cury := q[0][1]

        if isExit(curx, cury, entrance, n, m) {
            return dist[curx][cury]
        }

        q = q[1 : len(q)]

        newPositions := getNewPositions(maze, curx, cury, n, m)

        for _, pos := range newPositions {
            nx, ny := pos[0], pos[1]
            if dist[nx][ny] > 1 + dist[curx][cury] {
                dist[nx][ny] = 1 + dist[curx][cury]
                q = append(q, []int{nx, ny})
            }
        }
    }

    return -1
}

func getNewPositions(maze [][]byte, curx, cury, n, m int) [][]int {
    result := [][]int{}

    for i := 0; i < 4; i++ {
        nx := curx + dx[i]
        ny := cury + dy[i]

        if isIn(nx, ny, n, m) && maze[nx][ny] == '.' {
            result = append(result, []int{nx, ny})
        }
    }

    return result
}

func isIn(x, y, n, m int) bool {
    return 0 <= x && x < n && 0 <= y && y < m
}

func isExit(curx, cury int, entrance []int, n int, m int) bool {
    if curx == entrance[0] && cury == entrance[1] {
        return false
    }

    return curx == 0 || curx == n - 1 || cury == 0 || cury == m - 1
}
