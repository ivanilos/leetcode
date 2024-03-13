func pivotInteger(n int) int {
    pivot := 1

    leftSum := 0
    rightSum := ((1 + n) * n) / 2

    for pivot <= n {
        leftSum += pivot

        if leftSum == rightSum {
            return pivot
        }
        rightSum -= pivot
        pivot += 1
    }
    return -1
}
