const MAX_TIME = int64(1e7)
const MAX_TRIPS = int64(1e7)

func minimumTime(time []int, totalTrips int) int64 {
    maxPossibleTime := MAX_TIME * MAX_TRIPS

    minTime := int64(0)
    maxTime := maxPossibleTime
    result := int64(0)
    for minTime <= maxTime {
        checkTime := (minTime + maxTime) / 2

        if canDo(checkTime, time, totalTrips) {
            result = checkTime
            maxTime = checkTime - 1
        } else {
            minTime = checkTime + 1
        }
    }

    return result
}

func canDo(checkTime int64, time []int, totalTrips int) bool {
    tripsDone := int64(0)

    for _, t := range time {
        tripsDone += checkTime / int64(t)
    }

    return tripsDone >= int64(totalTrips)
}
