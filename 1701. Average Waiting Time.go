func averageWaitingTime(customers [][]int) float64 {
    result := 0
    chefTimer := 0
    for _, customer := range(customers) {
        arrival := customer[0]
        timer := customer[1]

        endMealTimer := max(chefTimer + timer, arrival + timer)
        result += endMealTimer - arrival
        chefTimer = endMealTimer
    }
    return float64(result) / float64(len(customers))
}
