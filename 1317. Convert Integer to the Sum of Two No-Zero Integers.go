func getNoZeroIntegers(n int) []int {
    for i := 1; i <= n / 2; i++ {
        if nonZero(i) && nonZero(n - i) {
            return []int{i, n - i}
        }
    }
    return []int{}
}

func nonZero(x int) bool {
    for x > 0 {
        if x % 10 == 0 {
            return false
        }
        x /= 10
    }
    return true
}
