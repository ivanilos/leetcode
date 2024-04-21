func validPath(n int, edges [][]int, source int, destination int) bool {
    adjList := map[int][]int{}

    for _, edge := range edges {
        a := edge[0]
        b := edge[1]
        adjList[a] = append(adjList[a], b)
        adjList[b] = append(adjList[b], a)
    }

    visited := make([]bool, n)
    DFS(source, adjList, visited)

    return visited[destination]
}

func DFS(node int, adjList map[int][]int, visited []bool) {
    visited[node] = true

    for _, viz := range adjList[node] {
        if !visited[viz] {
            DFS(viz, adjList, visited)
        }
    }
}
