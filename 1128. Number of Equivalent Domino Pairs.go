func numEquivDominoPairs(dominoes [][]int) int {
    const MAX_VAL = 9
    eq := make([][]int, MAX_VAL + 1)
    for i := 1; i <= MAX_VAL; i++ {
        eq[i] = make([]int, MAX_VAL + 1)
    }

    for _, domino := range dominoes {
        a := min(domino[0], domino[1])
        b := max(domino[0], domino[1])

        eq[a][b]++
    }

    result := 0
    for i := 1; i <= MAX_VAL; i++ {
        for j := 1; j <= MAX_VAL; j++ {
            result += (eq[i][j] * (eq[i][j] - 1)) / 2
        }
    }
    return result
}
