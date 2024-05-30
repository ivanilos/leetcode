func countTriplets(arr []int) int {
    n := len(arr)
    prefix := make([]int, n + 1)

    prefix[0] = 0
    for i := 0; i < n; i++ {
        prefix[i + 1] = arr[i] ^ prefix[i]
    }

    result := 0
    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            for k := j; k <= n; k++ {
                a := prefix[j - 1] ^ prefix[i - 1]
                b := prefix[k] ^ prefix[j - 1]
                if a == b {
                    result++
                }
            }
        }
    }
    return result
}
