func canArrange(arr []int, k int) bool {
    modFreq := make([]int, k)

    for _, val := range arr {
        val = ((val % k) + k) %k
        modFreq[val % k]++
    }

    for i := 1; i <= k / 2; i++ {
        if modFreq[i] != modFreq[k - i] {
            return false
        }
    }

    if modFreq[0] % 2 != 0 || (k % 2 == 0 && modFreq[k / 2] % 2 != 0) {
        return false
    }
    return true
}
