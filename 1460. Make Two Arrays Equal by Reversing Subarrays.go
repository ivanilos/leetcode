func canBeEqual(target []int, arr []int) bool {
    netTotal := map[int]int{}

    for _, val := range(target) {
        netTotal[val]++
    }
    for _, val := range(arr) {
        netTotal[val]--
    }
    for _, val := range(netTotal) {
        if val != 0 {
            return false
        }
    }
    return true
}
