func increasingTriplet(nums []int) bool {
    const INF int64 = int64(1e15)
    smallestMin := INF
    smallestMid := INF

    for _, val := range nums {
        num := int64(val)
        if num > smallestMid {
            return true
        }
        if num > smallestMin && num < smallestMid {
            smallestMid = num
        }
        if num < smallestMin {
            smallestMin = num
        }
    }
    return false
}
