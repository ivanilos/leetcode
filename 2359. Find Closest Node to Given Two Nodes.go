var INF = int(1e9)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
    distance1 := getDistance(edges, node1);
    distance2 := getDistance(edges, node2);

    fmt.Println(distance1, distance2)

    result := -1
    resultDist := INF
    for i := 0; i < len(edges); i++ {
        curMaxDist := max(distance1[i], distance2[i])
        if curMaxDist < resultDist {
            result = i
            resultDist = curMaxDist
        }
    }

    return result
}

func getDistance(edges []int, start int) []int {
    distance := make([]int, len(edges))
    for i := 0; i < len(edges); i++ {
        distance[i] = INF
    }
    distance[start] = 0

    DFS(start, edges, distance)

    return distance
}

func DFS(node int, edges []int, distance []int) {
    next := edges[node]

    if next != -1 && distance[next] > 1 + distance[node] {
        distance[next] = 1 + distance[node]
        DFS(next, edges, distance)
    }
}
