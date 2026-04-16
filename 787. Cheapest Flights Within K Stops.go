const INF = int(1e9)

type Node struct {
    idx int
    times int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    adjList := buildGraph(flights)

    dist := make([][]int, n * (k + 2))
    done := make([][]bool, n * (k + 2))
    for i := 0; i < n * (k + 2); i++ {
        dist[i] = make([]int, k + 2)
        done[i] = make([]bool, k + 2)
        for j := 0; j <= k + 1; j++ {
            dist[i][j] = INF
        }
    }

    dist[src][0] = 0

    for it := 0; it <= k; it++ {
        for step := 0; step < n; step++ {
            curMinDist := INF
            curMinDistIdx := -1
            for i := 0; i < n; i++ {
                if !done[i][it] && dist[i][it] < curMinDist {
                    curMinDist = dist[i][it]
                    curMinDistIdx = i
                }
            }

            if curMinDist == INF {
                break
            }

            done[curMinDistIdx][it] = true

            for _, edge := range adjList[curMinDistIdx] {
                neighbor := edge[0]
                price := edge[1]

                if dist[curMinDistIdx][it] + price < dist[neighbor][it + 1] {
                    dist[neighbor][it + 1] = dist[curMinDistIdx][it] + price
                }
            }
        }
    }

    result := INF
    for i := 0; i <= k + 1; i++ {
        if dist[dst][i] < result {
            result = dist[dst][i]
        }
    }

    if result == INF {
        result = -1
    }
    return result
}

func buildGraph(flights [][]int) map[int][][]int {
    adjList := map[int][][]int{}

    for _, flight := range flights {
        a := flight[0]
        b := flight[1]
        price := flight[2]

        adjList[a] = append(adjList[a], []int{b, price})
    }

    return adjList
}
