func minSwaps(s string) int {
    result := 0

    balance := 0
    for _, c := range s {
        if c == '[' {
            balance++
        } else {
            if balance == 0 {
                result++
                balance++
            } else {
                balance--
            }
        }
    }
    return result 
}
