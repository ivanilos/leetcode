func findChampion(n int, edges [][]int) int {
    losesQt := map[int]int{}

    for _, edge := range edges {
        losesQt[edge[1]]++
    }

    winners := []int{}
    for i := 0; i < n; i++ {
        if _, ok := losesQt[i]; !ok {
            winners = append(winners, i)
        }
    }

    if len(winners) > 1 {
        return -1
    }
    return winners[0]
}
