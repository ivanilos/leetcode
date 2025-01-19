type Node struct {
    x int
    y int
    maxWater int
}

func trapRainWater(heightMap [][]int) int {
    const INF = int(1e9)
    dx := []int{1, 0, -1, 0}
    dy := []int{0, 1, 0, -1}

    q := []Node{}

    maxWaterHeight := make([][]int, len(heightMap))
    for i := 0; i < len(heightMap); i++ {
        maxWaterHeight[i] = make([]int, len(heightMap[i]))
        for j := 0; j < len(heightMap[i]); j++ {
            maxWaterHeight[i][j] = INF

            if i == 0 || i == len(heightMap) - 1 || j == 0 || j == len(heightMap[i]) - 1 {
                maxWaterHeight[i][j] =  heightMap[i][j]
                q = append(q, Node{i, j, heightMap[i][j]})
            }
        }
    }

    for len(q) > 0 {
        next := q[0]
        q = q[1 : len(q)]

        for i := 0; i < 4; i++ {
            nx := next.x + dx[i];
            ny := next.y + dy[i];

            if isIn(nx, ny, heightMap) {
                nextMaxWater := max(heightMap[nx][ny], next.maxWater)

                if maxWaterHeight[nx][ny] > nextMaxWater {
                    maxWaterHeight[nx][ny] = nextMaxWater
                    q = append(q, Node{nx, ny, nextMaxWater})
                }
            }
        }
    }

    return getMaxWaterVolume(heightMap, maxWaterHeight)
}

func isIn(x, y int, heightMap [][]int) bool {
    return 0 <= x && x < len(heightMap) && 0 <= y && y < len(heightMap[x])
}

func getMaxWaterVolume(heightMap, maxWaterHeight [][]int) int {
    result := 0
    for i := 0 ; i < len(heightMap); i++ {
        for j := 0; j < len(heightMap[i]); j++ {
            result += maxWaterHeight[i][j] - heightMap[i][j]
        }
    }

    return result
}
