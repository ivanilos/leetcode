func maxTwoEvents(events [][]int) int {
    sort.Slice(events, func(a, b int) bool {
        return events[a][0] < events[b][0]
    })


    bestSuf := getBestSuf(events)
    result := 0
    for i := 0; i < len(events); i++ {
        result = max(result, events[i][2] + solve(i, events, bestSuf))
    }

    return result
}

func getBestSuf(events [][]int) []int {
    result := make([]int, len(events))
    result[len(events) - 1] = events[len(events) - 1][2]

    for i := len(events) - 2; i >= 0; i-- {
        result[i] = max(events[i][2], result[i + 1])
    }
    return result
}

func solve(curIdx int, events [][]int, bestSuf []int) int {
    low := curIdx + 1
    high := len(events) - 1
    bestIdx := -1

    for low <= high {
        mid := (low + high) / 2

        if events[mid][0] > events[curIdx][1] {
            high = mid - 1
            bestIdx = mid
        } else {
            low = mid + 1
        }
    }

    if bestIdx == -1 {
        return 0
    }
    return bestSuf[bestIdx]
}
