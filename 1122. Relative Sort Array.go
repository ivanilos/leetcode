func relativeSortArray(arr1 []int, arr2 []int) []int {
    orderMap := map[int]int{}

    for i := 0; i < len(arr2); i++ {
        orderMap[arr2[i]] = i
    }

    sort.Slice(arr1, func(i, j int) bool {
        pos1, ok1 := orderMap[arr1[i]]
        pos2, ok2 := orderMap[arr1[j]]

        if ok1 && ok2 {
            return pos1 < pos2
        } else if !ok1 && ok2 {
            return false
        } else if ok1 && !ok2 {
            return true
        } else {
            return arr1[i] < arr1[j]
        }
    })

    return arr1
}
