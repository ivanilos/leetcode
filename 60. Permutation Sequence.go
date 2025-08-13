func getPermutation(n int, k int) string {
    result := []byte{}

    used := make([]bool, n + 1)
    fat := getFactorials(n)

    for i := n; i > 0; i-- {
        for nextNum := 1; nextNum <= n; nextNum++ {
            if used[nextNum] {
                continue
            }

            if k - fat[i - 1] <= 0 {
                used[nextNum] = true
                result = append(result, byte(nextNum + int('0')))
                break
            } else {
                k -= fat[i - 1]
            }
        }
    }

    return string(result)
}

func getFactorials(n int) []int {
    fat := make([]int, n + 1)
    fat[0] = 1
    for i := 1; i <= n; i++ {
        fat[i] = i * fat[i - 1]
    }

    return fat
}
