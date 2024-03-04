func bagOfTokensScore(tokens []int, power int) int {
    result := 0

    sort.Ints(tokens)

    firstAvailable := 0
    lastAvailable := len(tokens) - 1

    score := 0
    for firstAvailable <= lastAvailable {
        if power >= tokens[firstAvailable] {
            score += 1
            power -= tokens[firstAvailable]
            firstAvailable += 1
        } else if score >= 1 {
            power += tokens[lastAvailable]
            score -= 1
            lastAvailable -= 1
        } else {
            break
        }
        result = max(result, score)
    }

    return result
}
