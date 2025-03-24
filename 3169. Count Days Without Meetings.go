func countDays(days int, meetings [][]int) int {
    events := [][]int{}

    for _, meeting := range meetings {
        events = append(events, []int{meeting[0], 1})
        events = append(events, []int{meeting[1] + 1, -1})
    }

    slices.SortFunc(events, func(a []int, b []int) int {
        return cmp.Compare(a[0], b[0])
    })

    result := 0
    lastDay := 1

    curMeetings := 0
    for _, event := range events {
        curDay := event[0]
        if curMeetings == 0 {
            result += max(0, curDay - lastDay)
        }

        lastDay = curDay
        curMeetings += event[1] 
    }
    if curMeetings == 0 {
        result += max(0, days - lastDay + 1)
    }

    return result
}
