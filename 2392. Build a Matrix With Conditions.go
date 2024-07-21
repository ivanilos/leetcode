func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
    rowOrder, errRow := topSort(rowConditions, k)
    colOrder, errCol := topSort(colConditions, k)

    if errRow != nil || errCol != nil {
        return [][]int{}
    }

    colOrderPos := map[int]int{}
    for i := 0; i < k; i++ {
        colOrderPos[colOrder[i]] = i
    }

    result := make([][]int, k)
    for i := 0; i < k ; i++ {
        result[i] = make([]int, k)
    }

    for i := 0; i < k; i++ {
        val := rowOrder[i]
        result[i][colOrderPos[val]] = val
    }
    return result
}

func topSort(conditions [][]int, k int) ([]int, error) {
    deg := make([]int, k + 1)
    adjList := make([][]int, k + 1)

    for i := 0; i < len(conditions); i++ {
        a := conditions[i][0]
        b := conditions[i][1]

        deg[b]++
        adjList[a] = append(adjList[a], b)
    }

    stack := []int{}
    for i := 1; i <= k; i++ {
        if deg[i] == 0 {
            stack = append(stack, i)
        }
    }

    result := []int{}
    for len(stack) > 0 {
        next := stack[len(stack) - 1]
        stack = stack[0 : len(stack) - 1]

        result = append(result, next)

        for _, viz := range(adjList[next]) {
            deg[viz]--
            if deg[viz] == 0 {
                stack = append(stack, viz)
            }
        }
    }

    if len(result) < k {
        return []int{}, errors.New("cycle in conditions")
    }
    return result, nil
}
