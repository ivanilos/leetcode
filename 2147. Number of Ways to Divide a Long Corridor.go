const MOD = int(1e9 + 7)

func numberOfWays(corridor string) int {
    result := 1

    curSeat := 0
    plantsAfter2Seats := 0
    for _, thing := range corridor {
        if thing == 'S' {
            if curSeat >= 2 {
                result = (result * (plantsAfter2Seats + 1)) % MOD
                plantsAfter2Seats = 0
                curSeat = 0
            }
            curSeat++
        } else if curSeat == 2 {
            plantsAfter2Seats++
        }
    }
    if curSeat < 2 {
        result *= 0
    }

    return result
}
