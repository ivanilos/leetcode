func maxScoreSightseeingPair(values []int) int {
    result := 0

    bestToPair := 0

    for i, val := range values {
        result = max(result, bestToPair + val - i)

        bestToPair = max(bestToPair, val + i)
    }
    return result
}
