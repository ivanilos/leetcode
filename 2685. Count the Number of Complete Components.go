func countCompleteComponents(n int, edges [][]int) int {
    adjMat := make([][]int, n)
    for i := 0; i < n; i++ {
        adjMat[i] = make([]int, n)
    }

    for _, edge := range edges {
        adjMat[edge[0]][edge[1]] = 1
        adjMat[edge[1]][edge[0]] = 1
    }

    return solve(n, adjMat)
}

func solve(n int, adjMat [][]int) int {
    comp := make([]int, n)
    used := make([]bool, n)

    nComps := 0
    for i := 0; i < n; i++ {
        if !used[i] {
            DFS(i, nComps, n, comp, used, adjMat)
            nComps++
        }
    }

    nodesInComponent := make([][]int, nComps)
    for i := 0; i < n; i++ {
        nodesInComponent[comp[i]] = append(nodesInComponent[comp[i]], i)
    }

    result := getCompleteComponents(nodesInComponent, adjMat)

    return result
}

func DFS(node, curComp, n int, comp []int, used []bool, adjMat [][]int) {
    used[node] = true
    comp[node] = curComp

    for i := 0; i < n; i++ {
        if !used[i] && adjMat[node][i] == 1 {
            DFS(i, curComp, n, comp, used, adjMat)
        }
    }
}

func getCompleteComponents(nodesInComponent [][]int, adjMat [][]int) int {
    result := 0

    for _, nodes := range nodesInComponent {
        good := true
        for i := 0; i < len(nodes); i++ {
            for j := i + 1; j < len(nodes); j++ {
                if adjMat[nodes[i]][nodes[j]] == 0 {
                    good = false
                    i = len(nodes)
                    break
                }
            }
        }

        if good {
            result++
        }
    }

    return result
}
