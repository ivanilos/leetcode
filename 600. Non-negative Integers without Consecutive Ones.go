type state struct {
    pos int
    found bool
    smaller bool
    lastDigit int
}

func findIntegers(n int) int {
    return n + 1 - getWithConsecutivesOnes(n)
}

func getWithConsecutivesOnes(n int) int {
    binaryNum := getBinary(n)

    dp := map[state]int{}

    return solve(state{0, false, false, 0}, binaryNum, dp)
}

func getBinary(n int) []int {
    result := []int{}
    for n > 0 {
        result = append(result, n & 1)
        n /= 2
    }
    slices.Reverse(result)

    return result
}

func solve(curState state, binaryNum []int, dp map[state]int) int {
    // base cases
    if curState.pos >= len(binaryNum) {
        if curState.found == true {
            return 1
        } else {
            return 0
        }
    }

    if _, ok := dp[curState]; ok {
        return dp[curState]
    }

    // put 0 at pos
    nextState := state{curState.pos + 1, curState.found, curState.smaller, 0}
    if binaryNum[curState.pos] == 1 {
        nextState.smaller = true
    }
    dp[curState] = solve(nextState, binaryNum, dp)

    // put 1 at pos if possible
    nextState = state{curState.pos + 1, curState.found, curState.smaller, 1}
    if curState.lastDigit == 1 {
        nextState.found = true
    }
    if curState.smaller || binaryNum[curState.pos] == 1 {
        dp[curState] += solve(nextState, binaryNum, dp)
    }

    return dp[curState]
}
