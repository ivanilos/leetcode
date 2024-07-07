func numWaterBottles(numBottles int, numExchange int) int {
    result := 0

    empty := 0
    full := numBottles
    for full > 0 {
        result += full
        exchanged := (full + empty) / numExchange
        empty = (full + empty) - ((full + empty) / numExchange) * numExchange
        full = exchanged

        fmt.Println(full, empty, exchanged, result)
    }
    return result
}
