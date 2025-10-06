var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func swimInWater(grid [][]int) int {
    n := len(grid)

    low := grid[0][0]
    high := n * n
    best := high

    for low <= high {
        mid := (low + high) / 2
        if canReach(grid, mid) {
            best = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return best
}

func canReach(grid [][]int, maxAllowed int) bool {
    n := len(grid)

    visited := make([][]bool, n)
    for i := 0; i < n; i++ {
        visited[i] = make([]bool, n)
    }

    q := [][]int{[]int{0, 0}}
    visited[0][0] = true

    for len(q) > 0 {
        x := q[0][0]
        y := q[0][1]
        q = q[1 : len(q)]

        for i := 0; i < 4; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            if isIn(nx, ny, n) && !visited[nx][ny] && grid[nx][ny] <= maxAllowed {
                visited[nx][ny] = true
                q = append(q, []int{nx, ny})
            }
        }
    }

    return visited[n - 1][n - 1] == true
}

func isIn(x, y, n int) bool {
    return 0 <= x && x < n && 0 <= y && y < n
}
