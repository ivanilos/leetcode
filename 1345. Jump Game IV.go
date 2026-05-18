const INF = int(1e9)

func minJumps(arr []int) int {
    valToPos := map[int][]int{}

    for i, val := range arr {
        valToPos[val] = append(valToPos[val], i)
    }

    dist := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        dist[i] = INF
    }
    dist[0] = 0

    queue := []int{0}
    for len(queue) > 0 {
        next := queue[0]
        queue = queue[1 : len(queue)]

        if next + 1 < len(arr) && dist[next + 1] > 1 + dist[next] {
            dist[next + 1] = 1 + dist[next]
            queue = append(queue, next + 1)
        }

        if next - 1 >= 0 && dist[next - 1] > 1 + dist[next] {
            dist[next - 1] = 1 + dist[next]
            queue = append(queue, next - 1)
        }

        for _, neighbor := range valToPos[arr[next]] {
            if dist[neighbor] > 1 + dist[next] {
                dist[neighbor] = 1 + dist[next]
                queue = append(queue, neighbor)
            }
        }

        valToPos[arr[next]] = []int{}
    }

    return dist[len(arr) - 1]
}
