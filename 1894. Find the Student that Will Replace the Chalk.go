func chalkReplacer(chalk []int, k int) int {
    sum := 0
    for _, c := range chalk {
        sum += c
    }

    leftBeforeLastRound := k % sum
    left := leftBeforeLastRound
    for i := 0; i < len(chalk); i++ {
        if chalk[i] > left {
            return i
        }
        left -= chalk[i]
    }
    return 0
}
