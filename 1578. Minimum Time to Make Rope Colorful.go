func minCost(colors string, neededTime []int) int {
    last := '+'
    sum := 0
    maxi := 0

    result := 0
    for i, c := range colors {
        if c == last {
            sum += neededTime[i]
            maxi = max(maxi, neededTime[i])
        } else {
            result += sum - maxi
            
            maxi = neededTime[i]
            sum = neededTime[i]
        }

        last = c
    }

    result += sum - maxi

    return result
}
