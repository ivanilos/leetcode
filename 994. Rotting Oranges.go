var ROTTEN = 2
var FRESH = 1

var INF = int(1e9)

var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func orangesRotting(grid [][]int) int {
    rotten := getRotten(grid)

    return BFS(grid, rotten)
}

func getRotten(grid [][]int) [][]int {
    result := [][]int{}

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == ROTTEN {
                result = append(result, []int{i, j})
            }
        }
    }

    return result
}

func BFS(grid [][]int, rotten [][]int) int {
    timer := initRottenTimer(grid, rotten)

    updateRottenTimer(grid, rotten, timer)

    return getAnswer(grid, timer)
}

func initRottenTimer(grid [][]int, rotten [][]int) [][]int {
    timer := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        timer[i] = make([]int, len(grid[i]))
        for j := 0; j < len(grid[i]); j++ {
            timer[i][j] = INF

            if grid[i][j] == ROTTEN {
                timer[i][j] = 0
            }
        }
    }

    return timer
}

func updateRottenTimer(grid [][]int, rotten [][]int, timer [][]int) {
    for len(rotten) > 0 {
        curX := rotten[0][0]
        curY := rotten[0][1]

        rotten = rotten[1 : len(rotten)]

        for i := 0; i < 4; i++ {
            nextX := curX + dx[i]
            nextY := curY + dy[i]

            if isIn(grid, nextX, nextY) && grid[nextX][nextY] == FRESH && timer[nextX][nextY] > timer[curX][curY] {
                timer[nextX][nextY] = 1 + timer[curX][curY]
                rotten = append(rotten, []int{nextX, nextY})
            }
        }
    }
}

func isIn(grid [][]int, x, y int) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}

func getAnswer(grid, timer [][]int) int {
    result := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == FRESH && timer[i][j] == INF {
                return -1
            }
            if grid[i][j] == FRESH || grid[i][j] == ROTTEN {
                result = max(result, timer[i][j])
            }
        }
    }

    return result
}
