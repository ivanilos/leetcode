const INF = int(1e9)

func minimumDistance(nums []int) int {
    valToPos := map[int][]int{}

    for i, num := range nums {
        valToPos[num] = append(valToPos[num], i)
    }

    result := INF
    for _, idxs := range valToPos {
        for i := 2; i < len(idxs); i++ {
            result = min(result, (idxs[i] - idxs[i - 1]) + (idxs[i] - idxs[i - 2]) + (idxs[i - 1] - idxs[i - 2]))
        }
    }

    if result == INF {
        return -1
    }
    return result
}
