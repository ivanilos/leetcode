func areSimilar(mat [][]int, k int) bool {
    lcm := 1

    for i := 0; i < len(mat); i++ {
        lcm = getLCM(lcm, getPeriod(mat, i))
    }

    return k % lcm == 0
}

func getPeriod(mat [][]int, row int) int {
    dir := 1
    if row % 2 == 1 {
        dir = -1
    }

    cols := len(mat[row])
    for step := 1; step <= cols; step++ {
        good := true

        for i := 0; i < cols; i++ {
            nextIdx := (((i + step * dir) % cols) + cols) % cols
            if mat[row][i] != mat[row][nextIdx] {
                good = false
                break
            }
        }

        if good {
            return step
        }
    }

    return cols
}

func getLCM(a, b int) int {
    return (a / gcd(a, b)) * b
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
