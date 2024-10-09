func minAddToMakeValid(s string) int {
    result := 0

    balance := 0
    for _, c := range s {
        if c == '(' {
            balance++
        } else {
            balance--
        }

        if balance < 0 {
            result++
            balance++
        }
    }

    result += balance
    return result
}
