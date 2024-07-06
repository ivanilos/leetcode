func passThePillow(n int, time int) int {
    moves := time % (2 * n - 2)

    if moves <= n - 1 {
        return 1 + moves
    } else {
        return n - (moves - (n - 1))
    }
}
