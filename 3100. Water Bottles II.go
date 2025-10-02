func maxBottlesDrunk(numBottles int, numExchange int) int {
    result := 0
    empty := 0
    for numBottles + empty >= numExchange {
        result += numBottles
        empty += numBottles
        numBottles = 0

        if empty >= numExchange {
            empty -= numExchange
            numExchange++
            numBottles++
        }
    }

    return result + numBottles
}
