func largestPathValue(colors string, edges [][]int) int {
    n := len(colors)
    adjList := buildGraph(edges)

    if hasCycle(n, adjList) {
        return -1
    }

    colorMax := make([][]int, n)
    for i := 0; i < n; i++ {
        colorMax[i] = make([]int, 26)
    }

    visited := make([]bool, n)

    for i := 0; i < n; i++ {
        if !visited[i] {
            DFS(i, visited, colors, colorMax, adjList)
        }    
    }

    maxi := 0
    for i := 0; i < n; i++ {
        for j := 0; j < 26; j++ {
            maxi = max(maxi, colorMax[i][j])
        }
    }

    return maxi
}

func DFS(node int, visited []bool, colors string, colorMax [][]int, adjList map[int][]int) {
    visited[node] = true

    for _, viz := range adjList[node] {
        if !visited[viz] {
            DFS(viz, visited, colors, colorMax, adjList)
        }
        for j := 0; j < 26; j++ {
            colorMax[node][j] = max(colorMax[node][j], colorMax[viz][j])
        }
    }

    curColor := int(colors[node] - 'a')
    colorMax[node][curColor]++
}

func buildGraph(edges [][]int) map[int][]int {
    adjList := map[int][]int{}

    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
    }

    return adjList
}

func hasCycle(n int, adjList map[int][]int) bool {
    color := make([]int, n)

    hasCycle := false
    for i := 0; i < n; i++ {
        if color[i] == 0 {
            hasCycle = hasCycle || checkCycle(i, color, adjList)
        }
    }

    return hasCycle
}

func checkCycle(node int, color []int, adjList map[int][]int) bool {
    color[node] = 1

    hasCycle := false
    for _, viz := range adjList[node] {
        if color[viz] == 1 {
            return true
        } else if color[viz] == 0 {
            hasCycle = hasCycle || checkCycle(viz, color, adjList)
        }
    }

    color[node] = 2
    return hasCycle
}
