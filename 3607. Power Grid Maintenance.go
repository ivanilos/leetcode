func DFS(node, ncomp int, nodeToCompMap map[int]int, visited []bool, componentNodes [][]int, connectionsMap map[int][]int) {
    visited[node] = true

    nodeToCompMap[node] = ncomp

    componentNodes[ncomp] = append(componentNodes[ncomp], node)

    for _, viz := range connectionsMap[node] {
        if !visited[viz] {
            DFS(viz, ncomp, nodeToCompMap, visited, componentNodes, connectionsMap)
        }
    }
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
    connectionsMap := map[int][]int{}
    for _, connection := range connections {
        a := connection[0]
        b := connection[1]

        connectionsMap[a] = append(connectionsMap[a], b)
        connectionsMap[b] = append(connectionsMap[b], a)
    }

    visited := make([]bool, c + 1)
    ncomp := 0
    componentNodes := [][]int{}
    nodeToCompMap := map[int]int{}
    for i := 1; i <= c; i++ {
        if !visited[i] {
            componentNodes = append(componentNodes, []int{})
            DFS(i, ncomp, nodeToCompMap, visited, componentNodes, connectionsMap)

            ncomp++
        }
    }

    nextOnlineIdx := make([]int, ncomp)
    for i := 0; i < len(componentNodes); i++ {
        slices.Sort(componentNodes[i])
    }

    result := []int{}
    offline := make([]bool, c + 1)
    for _, query := range queries {
        op := query[0]
        node := query[1]

        if op == 1 {
            comp := nodeToCompMap[node]

            if !offline[node] {
                result = append(result, node)
            } else {
                curQueryResult := -1
                for nextOnlineIdx[comp] < len(componentNodes[comp]) {
                    if !offline[componentNodes[comp][nextOnlineIdx[comp]]] {
                        curQueryResult = componentNodes[comp][nextOnlineIdx[comp]]
                        break
                    }
                    nextOnlineIdx[comp]++
                }
                result = append(result, curQueryResult)
            }
        } else {
            offline[node] = true
        }
    }

    return result
}
