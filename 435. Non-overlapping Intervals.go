func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][1] < intervals[b][1]
    })

    lastPoint := intervals[0][0]
    result := 0

    for _, interval := range intervals {
        if interval[0] >= lastPoint {
            result++
            lastPoint = interval[1]
        }
    }

    return len(intervals) - result
}
