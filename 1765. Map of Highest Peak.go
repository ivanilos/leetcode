type node struct {
    x int
    y int
}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func highestPeak(isWater [][]int) [][]int {
    result := make([][]int, len(isWater))
    for i := 0; i < len(isWater); i++ {
        result[i] = make([]int, len(isWater[i]))
        for j := 0; j < len(isWater[i]); j++ {
            result[i][j] = -1
        }
    }


    q := []node{}
    for i := 0; i < len(isWater); i++ {
        for j := 0; j < len(isWater[i]); j++ {
            if isWater[i][j] == 1 {
                q = append(q, node{i, j})
                result[i][j] = 0
            }
        }
    }

    for len(q) > 0 {
        nextNode := q[0]
        x := nextNode.x
        y := nextNode.y

        q = q[1 : len(q)]

        for i := 0; i < 4; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            if isIn(nx, ny, isWater) && result[nx][ny] == -1 {
                result[nx][ny] = 1 + result[x][y]
                q = append(q, node{nx, ny})
            }
        } 
    }

    return result
}

func isIn(x, y int, grid [][]int) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}
