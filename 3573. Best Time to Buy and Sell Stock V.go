type node struct {
    pos int
    opLeft int
    stockState int
}

const DONT_HAVE = 0
const HAS = 1
const OWN = 2

const INF = int64(1e16)

func maximumProfit(prices []int, k int) int64 {
    dp := map[node]int64{}

    start := node{0, 2 * k, 0}
    result := solve(start, prices, dp)

    return result
}

func solve(state node, prices []int, dp map[node]int64) int64 {
    pos := state.pos
    opLeft := state.opLeft
    stockState := state.stockState

    if pos >= len(prices) {
        if stockState != OWN {
            return 0
        }
        return -INF
    }

    if _, ok := dp[state]; ok {
        return dp[state]
    }

    // do nothing
    newState := node{pos + 1, opLeft, stockState}
    result := solve(newState, prices, dp)

    // buy stock
    if opLeft > 0 && (stockState == DONT_HAVE || stockState == OWN) {
        newState := node{pos + 1, opLeft - 1, HAS}
        if stockState == OWN {
            newState.stockState = DONT_HAVE
        }
        result = max(result, -int64(prices[pos]) + solve(newState, prices, dp))
    }

    // sell stock
    if opLeft > 0 && (stockState == DONT_HAVE || stockState == HAS) {
        newState := node{pos + 1, opLeft - 1, OWN}
        if stockState == HAS {
            newState.stockState = DONT_HAVE
        }
        result = max(result, int64(prices[pos]) + solve(newState, prices, dp))
    }

    dp[state] = result
    return result
}
