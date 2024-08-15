func lemonadeChange(bills []int) bool {
    got := []int{0, 0}

    for _, bill := range(bills) {
        if bill == 5 {
            got[0]++
        } else if bill == 10 {
            if got[0] > 0 {
                got[0]--
                got[1]++
            } else {
                return false
            }
        } else {
            if got[0] > 0 && got[1] > 0 {
                got[0]--
                got[1]--
            } else if got[0] >= 3 {
                got[0] -= 3
            } else {
                return false
            }
        }
    }
    return true
}
