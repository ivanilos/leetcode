func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    dist := getDistMatrix(n, edges)

    return solve(n, dist, distanceThreshold)
}

func getDistMatrix(n int, edges [][]int) [][]int {
    INF := int(1e9)

    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dist[i][j] = INF
        }
        dist[i][i] = 0
    }

    for _, edge := range(edges) {
        from := edge[0]
        to := edge[1]
        weight := edge[2]

        dist[from][to] = min(dist[from][to], weight)
        dist[to][from] = min(dist[to][from], weight)
    }

    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
            }
        }
    }
    return dist
}

func solve(n int, dist [][]int, distanceThreshold int) int {
    result := n
    best := n + 1
    for i := n - 1; i >= 0; i-- {
        total := 0
        for j := 0; j < n; j++ {
            if dist[i][j] <= distanceThreshold {
                total++
            }
        }

        if total < best {
            best = total
            result = i
        }
    }
    return result
}
