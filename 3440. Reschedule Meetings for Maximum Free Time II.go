func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    startTime = append([]int{0}, startTime...)
    startTime = append(startTime, eventTime)
    endTime = append([]int{0}, endTime...)
    endTime = append(endTime, eventTime)

    prefSpace := make([]int, len(startTime))
    maxPrefSpace := make([]int, len(startTime))
    maxPrefSpace[0] = 0
    for i := 1; i < len(startTime); i++ {
        prefSpace[i] = startTime[i] - endTime[i - 1]
        maxPrefSpace[i] = max(maxPrefSpace[i - 1], prefSpace[i])
    }

    sufSpace := make([]int, len(endTime))
    maxSufSpace := make([]int, len(endTime))
    maxSufSpace[len(endTime) - 1] = 0
    for i := len(endTime) - 2; i >= 0; i-- {
        sufSpace[i] = startTime[i + 1] - endTime[i]
        maxSufSpace[i] = max(maxSufSpace[i + 1], sufSpace[i])
    }

    result := 0
    for i := 1; i < len(startTime) - 1; i++ {
        duration := (endTime[i] - startTime[i])
        if max(maxPrefSpace[i - 1], maxSufSpace[i + 1]) >= duration {
            result = max(result, startTime[i + 1] - endTime[i - 1])
        } else {
            result = max(result, startTime[i + 1] - endTime[i - 1] - duration)
        }
    }

    return result
}
