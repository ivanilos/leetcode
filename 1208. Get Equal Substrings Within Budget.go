func equalSubstring(s string, t string, maxCost int) int {
    cost := make([]int, len(s))

    for i := 0; i < len(s); i++ {
        cost[i] = abs(int(s[i]) - int(t[i]))
    }

    return solve(cost, maxCost)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func solve(cost []int, maxCost int) int {
    left := 0
    
    curCost := 0
    best := 0
    for right := 0; right < len(cost); right++ {
        curCost += cost[right]

        for curCost > maxCost {
            curCost -= cost[left]
            left++
        }
        best = max(best, right - left + 1)
    }
    return best
}
