type Edge struct {
    neighbour int
    dir int
}

func minReorder(n int, connections [][]int) int {
    out := make([]int, n)

    adjList := make([][]Edge, n)

    for _, edge := range connections {
        u := edge[0]
        v := edge[1]

        out[u]++

        adjList[u] = append(adjList[u], Edge{v, 1})
        adjList[v] = append(adjList[v], Edge{u, -1})
    }

    result := DFS(0, -1, out, adjList)

    return result
}

func DFS(node, p int, out []int, adjList [][]Edge) int {
    result := out[node]
    for _, e := range adjList[node] {
        if e.neighbour != p {
            if e.dir == -1 {
                out[e.neighbour]--
            }
            result += DFS(e.neighbour, node, out, adjList)
        }
    }

    return result
}
