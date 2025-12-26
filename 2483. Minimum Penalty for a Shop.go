func bestClosingTime(customers string) int {
    sufY := 0

    for _, customer := range customers {
        if customer == 'Y' {
            sufY++
        }
    }

    result := 0
    minCost := sufY

    prefN := 0
    for i, customer := range customers {
        if customer == 'N' {
            prefN++
        } else {
            sufY--
        }

        newCost := sufY + prefN
        if newCost < minCost {
            minCost = newCost
            result = i + 1
        }
    }

    return result
}
