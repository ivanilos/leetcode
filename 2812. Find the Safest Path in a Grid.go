type pair struct {
    x, y int
}

func maximumSafenessFactor(grid [][]int) int {
    n := len(grid)

    thieves := []pair{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                thieves = append(thieves, pair{i, j})
            }
        }
    }

    dist := calcThiefDist(n, thieves)

    for result := n * n; result >= 0; result-- {
        if BFS(n, dist, result) {
            return result
        }
    }
    return 0
}

func calcThiefDist(n int, thieves []pair) [][]int {
    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, n)

        for j := 0; j < n; j++ {
            dist[i][j] = 2 * n + 1
        }
    }

    for _, thief := range thieves {
        x := thief.x
        y := thief.y
        dist[x][y] = 0
    }

    dx := []int{1, 0, -1, 0}
    dy := []int{0, -1, 0, 1}

    for len(thieves) > 0 {
        x := thieves[0].x
        y := thieves[0].y
        thieves = thieves[1:len(thieves)]

        for i := 0; i < 4; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            if isIn(nx, ny, n) && dist[nx][ny] > 1 + dist[x][y] {
                dist[nx][ny] = 1 + dist[x][y]
                thieves = append(thieves, pair{nx, ny})
            }
        }
    }
    return dist
}

func BFS(n int, thievesDist[][]int, minDistAllowed int) bool {
    if thievesDist[0][0] < minDistAllowed {
        return false
    }

    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, n)
    }

    dx := []int{1, 0, -1, 0}
    dy := []int{0, -1, 0, 1}

    visited[0][0] = true
    queue := []pair{pair{0, 0}}
    for len(queue) > 0 {
        x := queue[0].x
        y := queue[0].y
        queue = queue[1:len(queue)]

        for i := 0; i < 4; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            if isIn(nx, ny, n) && !visited[nx][ny] && thievesDist[nx][ny] >= minDistAllowed {
                visited[nx][ny] = true
                queue = append(queue, pair{nx, ny})
            }
        }
    }
    return visited[n - 1][n - 1]
}

func isIn(x, y, n int) bool {
    return 0 <= x && x < n && 0 <= y && y < n
}
