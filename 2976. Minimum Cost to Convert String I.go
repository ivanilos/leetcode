const ALPHA int = 26
const INF int64 = int64(1e15)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    dist := getDistMatrix(original, changed, cost)

    result := solve(dist, source, target)
    return result
}

func getDistMatrix(original []byte, changed []byte, cost []int) [][]int64 {
    dist := make([][]int64, ALPHA)
    for i := 0; i < ALPHA; i++ {
        dist[i] = make([]int64, ALPHA)
        for j := 0; j < ALPHA; j++ {
            dist[i][j] = INF
        }
        dist[i][i] = 0
    }

    for i := 0; i < len(original); i++ {
        from := original[i] - 'a'
        to := changed[i] - 'a'
        weight := int64(cost[i])

        dist[from][to] = min(dist[from][to], weight) 
    }
    return dist
}

func solve(dist [][]int64, source string, target string) int64 {
    for k := 0; k < ALPHA; k++ {
        for i := 0; i < ALPHA; i++ {
            for j := 0; j < ALPHA; j++ {
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
            }
        }
    }

    result := int64(0)
    for i := 0; i < len(source); i++ {
        from := source[i] - 'a'
        to := target[i] - 'a'

        if dist[from][to] == INF {
            return -1
        } else {
            result += dist[from][to]
        }
    }
    return result
}
