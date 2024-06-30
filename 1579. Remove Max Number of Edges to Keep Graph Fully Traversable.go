func maxNumEdgesToRemove(n int, edges [][]int) int {
    sharedEdges := [][]int{}
    aliceEdges := [][]int{}
    bobEdges := [][]int{}

    for _, edge := range(edges) {
        if edge[0] == 1 {
            aliceEdges = append(aliceEdges, []int{edge[1], edge[2]})
        } else if edge[0] == 2 {
            bobEdges = append(bobEdges, []int{edge[1], edge[2]})
        } else {
            sharedEdges = append(sharedEdges, []int{edge[1], edge[2]})
        }
    }

    aliceSet := buildSet(n)
    bobSet := buildSet(n)

    aliceResult := 0
    bobResult := 0
    sharedResult := 0
    for _, edge := range(sharedEdges) {
        sharedResult += unite(edge[0], edge[1], aliceSet)
        unite(edge[0], edge[1], bobSet)
    }

    for _, edge := range(aliceEdges) {
        aliceResult += unite(edge[0], edge[1], aliceSet)
    }

    for _, edge := range(bobEdges) {
        bobResult += unite(edge[0], edge[1], bobSet)
    }

    if sharedResult + aliceResult < n - 1 || sharedResult + bobResult < n - 1 {
        return -1
    }

    return len(edges) - (sharedResult + aliceResult + bobResult)
}

func find(i int, s []int) int {
    if i == s[i] {
        return i
    }
    s[i] = find(s[i], s)
    return s[i]
}

func unite(a, b int, s []int) int {
    a = find(s[a], s)
    b = find(s[b], s)

    if a == b {
        return 0
    }

    s[a] = s[b]
    return 1
}

func buildSet(n int) []int {
    s := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        s[i] = i
    }
    return s
}
