func findCircleNum(isConnected [][]int) int {
    nodes := len(isConnected)
    visited := make([]bool, nodes)

    result := 0
    for i := 0; i < nodes; i++ {
        if !visited[i] {
            result++
            DFS(i, isConnected, visited)
        }
    }

    return result
}

func DFS(node int, isConnected [][]int, visited []bool) {
    visited[node] = true

    for city := 0; city < len(isConnected); city++ {
        if isConnected[node][city] == 1 && !visited[city] {
            DFS(city, isConnected, visited)
        }
    }
}
