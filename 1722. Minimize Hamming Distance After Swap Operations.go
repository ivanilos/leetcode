func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    swaps := map[int][]int{}

    for _, swapPair := range allowedSwaps {
        a := swapPair[0]
        b := swapPair[1]

        swaps[a] = append(swaps[a], b)
        swaps[b] = append(swaps[b], a)
    }

    components := [][]int{}
    used := map[int]bool{}
    for idx := 0; idx < len(source); idx++ {
        if _, ok := used[idx]; !ok {
            newComponent := []int{}
            DFS(idx, used, swaps, &newComponent)

            components = append(components, newComponent)
        }
    }

    result := 0
    for _, component := range components {
        freqSource := map[int]int{}
        freqTarget := map[int]int{}

        for _, idx := range component {
            freqSource[source[idx]]++
            freqTarget[target[idx]]++
        }

        for key, _ := range freqSource {
            result += max(0, freqSource[key] - freqTarget[key])
        }
    }

    return result
}

func DFS(node int, used map[int]bool, adjList map[int][]int, curComponent *[]int) {
    used[node] = true

    *curComponent = append(*curComponent, node)

    for _, neighbor := range adjList[node] {
        if _, ok := used[neighbor]; !ok {
            DFS(neighbor, used, adjList, curComponent)
        }
    }
}
