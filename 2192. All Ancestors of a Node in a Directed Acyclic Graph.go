func getAncestors(n int, edges [][]int) [][]int {
    adjList := make([][]int, n)

    for _, edge := range(edges) {
        a := edge[0]
        b := edge[1]
        adjList[a] = append(adjList[a], b)
    }

    result := make([][]int, n)
    visited := make([]bool, n)
    for i := 0; i < n; i++ {
        for i := 0; i < n; i++ {
            visited[i] = false
        }
        DFS(i, i, visited, adjList, result)
    }
    for i := 0; i < n; i++ {
        slices.Sort(result[i])
    }
    return result
}

func DFS(node, start int, visited []bool, adjList, result [][]int) {
    visited[node] = true
    if node != start {
        result[node] = append(result[node], start)
    }

    for _, viz := range(adjList[node]) {
        if !visited[viz] {
            DFS(viz, start, visited, adjList, result)
        }
    }
}
