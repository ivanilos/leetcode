type Event struct {
    time int
    val int
}

func minGroups(intervals [][]int) int {
    events := []Event{}

    for _, interval := range intervals {
        left := interval[0]
        right := interval[1]

        events = append(events, Event{left, 1})
        events = append(events, Event{right, -1})
    }

    sort.Slice(events, func(a, b int) bool {
        return events[a].time < events[b].time || events[a].time == events[b].time && events[a].val > events[b].val
    })

    result := 0
    cur := 0

    for _, event := range events {
        cur += event.val
        result = max(result, cur)
    }
    return result
}
