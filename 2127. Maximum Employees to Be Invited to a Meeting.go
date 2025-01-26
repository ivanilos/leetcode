func maximumInvitations(favorite []int) int {
    indegree := getIndegree(favorite)

    q := getWithoutIndegree(indegree)

    maxDepth := make([]int, len(favorite))
    for len(q) > 0 {
        next := q[0]
        q = q[1 : len(q)]

        fmt.Println(next)

        maxDepth[favorite[next]] = max(maxDepth[favorite[next]], 1 + maxDepth[next])

        indegree[favorite[next]]--
        if indegree[favorite[next]] == 0 {
            q = append(q, favorite[next])
        }
    }

    result := solve(favorite, indegree, maxDepth)
    return result
}

func getIndegree(favorite []int) []int {
    indegree := make([]int, len(favorite))

    for _, fav := range favorite {
        indegree[fav]++
    }
    return indegree
}

func getWithoutIndegree(indegree []int) []int {
    result := []int{}
    for i := 0; i < len(indegree); i++ {
        if indegree[i] == 0 {
            result = append(result, i)
        }
    }

    return result;
}

func solve(favorite, indegree, maxDepth []int) int {
    maxCycle := 0
    maxWithCyclesOfLen2 := 0
    for i := 0; i < len(favorite); i++ {
        if indegree[i] != 0 {
            cycleLen := 0
            next := i
            for indegree[next] != 0 {
                indegree[next]--
                cycleLen++
                next = favorite[next]
            }

            maxCycle = max(maxCycle,cycleLen)

            if cycleLen == 2 {
                maxWithCyclesOfLen2 += maxDepth[next] + 1 + maxDepth[favorite[next]] + 1
            }

        }
    }

    return max(maxCycle, maxWithCyclesOfLen2)
}
