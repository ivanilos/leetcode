func smallestRepunitDivByK(k int) int {
    result := 1
    found := map[int]bool{}

    cur := 1 % k
    for cur % k != 0 {
        if _, ok := found[cur]; ok {
            return -1
        }
        found[cur] = true

        cur = (10 * cur + 1) % k
        result++
    }

    return result
}
