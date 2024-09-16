func findMinDifference(timePoints []string) int {
    timesInMinutes := make([]int, len(timePoints))

    for i := 0; i < len(timePoints); i++ {
        timesInMinutes[i] = convertToMinutes(timePoints[i])
    }

    sort.Ints(timesInMinutes)

    result := 24 * 60 - (timesInMinutes[len(timesInMinutes) - 1] - timesInMinutes[0])
    for i := 1; i < len(timesInMinutes); i++ {
        result = min(result, timesInMinutes[i] - timesInMinutes[i - 1])
    }
    return result
}

func convertToMinutes(timePoint string) int {
    hours, _ := strconv.Atoi(timePoint[0:2])
    minutes, _ := strconv.Atoi(timePoint[3:5])

    return hours * 60 + minutes
}
