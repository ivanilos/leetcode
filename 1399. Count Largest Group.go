func countLargestGroup(n int) int {
    freq := map[int]int{}

    for i := 1; i <= n; i++ {
        s := getSum(i)
        freq[s]++
    }

    maxi := 0
    qt := 0
    for _, val := range freq {
        if val == maxi {
            qt++
        } else if val > maxi {
            qt = 1
            maxi = val
        }
    }
    return qt
}

func getSum(x int) int {
    result := 0

    for x > 0 {
        result += x % 10
        x /= 10
    }
    return result
}
