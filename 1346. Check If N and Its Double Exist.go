func checkIfExist(arr []int) bool {
    found := map[int]bool{}

    sort.Slice(arr, func(a, b int) bool {
        return abs(arr[a]) > abs(arr[b])
    })
    
    for _, val := range arr {
        if _, ok := found[2 * val]; ok {
            return true
        }
        found[val] = true
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
