type Event struct {
    timer int
    isPerson bool
    personIdx int
    isFlowerStart bool
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
    events := getOrderedEvents(flowers, people)

    result := make([]int, len(people))
    curBloom := 0
    for _, event := range events {
        if event.isPerson {
            result[event.personIdx] = curBloom
        } else if event.isFlowerStart {
            curBloom++
        } else {
            curBloom--
        }
    }

    return result
}

func getOrderedEvents(flowers [][]int, people []int) []Event {
    events := []Event{}

    for _, flower := range flowers {
        events = append(events, Event{flower[0], false, -1, true})
        events = append(events, Event{flower[1] + 1, false, -1, false})
    }

    for i, person := range people {
        events = append(events, Event{person, true, i, false})
    }

    slices.SortFunc(events, func(a, b Event) int {
        if a.timer == b.timer {
            if a.isPerson {
                return 1
            } else {
                return -1
            }
        }

        return a.timer - b.timer
    })

    return events
}
