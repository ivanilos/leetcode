func findOrder(numCourses int, prerequisites [][]int) []int {
    adjList := map[int][]int{}
    indegree := make([]int, numCourses)

    for _, prerequisite := range prerequisites {
        a, b := prerequisite[0], prerequisite[1]

        indegree[a]++
        adjList[b] = append(adjList[b], a)
    }

    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    result := []int{}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        result = append(result, cur)

        for _, neighbor := range adjList[cur] {
            indegree[neighbor]--

            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    if len(result) < numCourses {
        return []int{}
    }
    return result
}
