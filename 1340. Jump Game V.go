func maxJumps(arr []int, d int) int {
    adjList, indegree := getGraphAndIndegree(arr, d)

    return solve(adjList, indegree)
}

func solve(adjList [][]int, indegree []int) int {
    n := len(indegree)
    queue := []int{}

    for i := 0; i < n; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    dist := make([]int, n)

    result := 0
    for len(queue) > 0 {
        next := queue[0]
        queue = queue[1:]

        result = max(result, dist[next])

        for _, neighbor := range adjList[next] {
            dist[neighbor] = 1 + dist[next]
            indegree[neighbor]--

            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    return result + 1
}

func getGraphAndIndegree(arr []int, d int) ([][]int, []int) {
    adjList := make([][]int, len(arr))
    indegree := make([]int, len(arr))

    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr) && j <= i + d; j++ {
            if arr[i] > arr[j] {
                adjList[i] = append(adjList[i], j)
                indegree[j]++
            } else {
                break
            }
        }
    }

    for i := len(arr) - 1; i >= 0; i-- {
        for j := i - 1; j >= 0 && j >= i - d; j-- {
            if arr[i] > arr[j] {
                adjList[i] = append(adjList[i], j)
                indegree[j]++
            } else {
                break
            }
        }
    }

    return adjList, indegree
}
