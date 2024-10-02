func arrayRankTransform(arr []int) []int {
    rankMap := getRankMap(arr)

    for i := 0; i < len(arr); i++ {
        arr[i] = rankMap[arr[i]]
    }
    return arr
}

func getRankMap(arr []int) map[int]int {
    val := make([]int, len(arr))
    copy(val, arr)

    sort.Ints(val)

    result := map[int]int{}
    nextRank := 1
    for i := 0; i < len(val); i++ {
        if _, ok := result[val[i]]; !ok {
            result[val[i]] = nextRank
            nextRank++
        }
    }
    return result
}
