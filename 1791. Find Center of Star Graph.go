func findCenter(edges [][]int) int {
    adjCount := map[int]int{}

    adjCount[edges[0][0]]++
    adjCount[edges[0][1]]++
    adjCount[edges[1][0]]++
    adjCount[edges[1][1]]++

    for key, val := range(adjCount) {
        if val >= 2 {
            return key
        }
    }
    panic("should not reach this")
}
