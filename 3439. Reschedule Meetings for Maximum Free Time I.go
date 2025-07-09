func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    const INF = int(1e9)

    startTime = append(startTime, eventTime)
    endTime = append(endTime, INF)

    totalDuration := 0
    for i := 0; i < k; i++ {
        totalDuration += endTime[i] - startTime[i]
    }

    result := 0
    l := 0
    for r := k; r < len(startTime); r++ {
        if r == k {
            if totalDuration < startTime[r] {
                result = max(result, startTime[r] - totalDuration)
            }
        } else {
            if totalDuration < startTime[r] - endTime[l - 1] {
                result = max(result, (startTime[r] - endTime[l - 1]) - totalDuration)
            }
        }

        totalDuration += endTime[r] - startTime[r]
        totalDuration -= endTime[l] - startTime[l]
        l++
    }

    return result
}
